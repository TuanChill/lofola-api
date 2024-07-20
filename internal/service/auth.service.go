package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/configs/common/constants"
	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/models"
	"github.com/tuanchill/lofola-api/internal/repo"
	"github.com/tuanchill/lofola-api/internal/repo/redis"
	"github.com/tuanchill/lofola-api/pkg/helpers"
	mailer "github.com/tuanchill/lofola-api/pkg/helpers/mail"
	"github.com/tuanchill/lofola-api/pkg/logger"
	"github.com/tuanchill/lofola-api/pkg/response"
	"github.com/tuanchill/lofola-api/pkg/utils"
	"gorm.io/gorm"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// Register handles the registration process for a user.
// It checks for user spam, validates the request body, checks if the user already exists,
// creates a new user if not, generates a verification link, sends an email for verification,
// and returns the registration response containing the user ID, email, and verification token.
// If any error occurs during the process, it returns an appropriate error response.
//
// @Summary Register a new user
// @Description Handles the registration process for a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.BodyRegisterRequest true "Registration request body"
// @Param X-Device-Id header string true "Device ID"
// @Success 200 {object} models.RegistrationResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/register [post]
func (a *AuthService) Register(c *gin.Context) *models.UserResponseBody {
	// get data from body
	reqBody := models.UserRequestBody{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.BadRequestErrorWithFields(c, response.ErrCodeValidation, utils.GetObjMessage(err))
		return nil
	}

	// check spam user
	// if spamResponse := redis.SpamUser(c, global.RDB, constants.SpamKey, constants.RequestThreshold); spamResponse != nil {
	// 	if spamResponse.IsSpam {
	// 		ttl := fmt.Sprintf("You are blocked for %d seconds", spamResponse.ExpireTime)
	// 		response.BadRequestError(c, response.ErrIpBlackList, ttl)
	// 		return nil
	// 	}
	// }

	// get user by email
	user, err := repo.GetDetailUserByEmail(global.MDB, reqBody.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			response.BadRequestError(c, response.ErrCodeResourceConflict, err.Error())
			return nil
		}
	}

	// if user exists
	if user.ID > 0 {
		response.BadRequestError(c, response.ErrCodeResourceConflict, "Email already exists")
		return nil
	}

	// get user by username
	oldUser, err := repo.GetDetailUserByUsername(global.MDB, reqBody.UserName)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			response.BadRequestError(c, response.ErrCodeResourceConflict, err.Error())
			return nil
		}
	}

	// if user exists
	if oldUser.ID > 0 {
		response.BadRequestError(c, response.ErrCodeResourceConflict, "UserName already exists")
		return nil
	}

	// hash password
	hashedPassword, err := helpers.HashPassword(reqBody.Password)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer)
		return nil
	}

	fmt.Println("hashedPassword", hashedPassword)

	reqBody.Password = hashedPassword

	// create user
	newUser, err := repo.CreateUser(global.MDB, reqBody)
	if err != nil {
		response.InternalServerError(c, response.ErrCodeDBQuery)
		return nil
	}

	// send email verification asynchronously
	go func(user models.User) {
		otp := helpers.GenerateOTP(6)
		otpMailData := models.DataOtpMail{
			Title:  "OTP Verification",
			To:     user.Email,
			Expire: helpers.FormatDuration(constants.ExpiresOTP),
			Otp:    otp,
			Name:   user.UserName,
		}

		if err := mailer.SendOptMail(otpMailData); err != nil {
			// Log error if needed, but do not block the main process
			logger.LogError(fmt.Sprintf("Failed to send OTP email to %s: %v\n", user.Email, err))
		}
	}(newUser)

	return &models.UserResponseBody{
		UserName: newUser.UserName,
		Email:    newUser.Email,
	}
}

func (a *AuthService) Login(c *gin.Context) *models.LoginResponse {
	var reqBody *models.BodyLoginRequest

	// get and validate
	if err := c.ShouldBindBodyWithJSON(&reqBody); err != nil {
		response.BadRequestErrorWithFields(c, response.ErrCodeValidation, utils.GetObjMessage(err))
		return nil
	}

	//  check identifier is email or username
	caseIdentifier := helpers.CheckIdentifier(reqBody.Identifier)

	var user models.User // Giả sử User là kiểu dữ liệu của người dùng, thay bằng kiểu dữ liệu thực tế của bạn
	var err error

	switch caseIdentifier {
	case "email":
		// get user by email
		user, err = repo.GetDetailUserByEmail(global.MDB, reqBody.Identifier)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				response.BadRequestError(c, response.ErrCodeResourceConflict, err.Error())
				return nil
			}
		}
	case "username":
		// get user by username
		user, err = repo.GetDetailUserByUsername(global.MDB, reqBody.Identifier)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				response.BadRequestError(c, response.ErrCodeResourceConflict, err.Error())
				return nil
			}
		}
	}

	// if user not found
	if user.ID == 0 {
		// Handle case where user is not found
		response.BadRequestError(c, response.ErrCodeResourceConflict, "User not found")
		return nil
	}

	// check password
	if err := helpers.ComparePassword(user.Password, reqBody.Password); err != nil {
		response.BadRequestError(c, response.ErrCodeResourceConflict, "Password is incorrect")
		return nil
	}

	// check user is active
	if !user.IsActive {
		response.ForbiddenError(c, response.ErrCodeForbidden, "User is not active")
		return nil
	}

	// generate token
	accessToken, refreshToken, err := helpers.JwtProvider(models.PayloadToken{
		ID:       int(user.ID),
		Email:    user.Email,
		UserName: user.UserName,
	})
	if err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer, "Failed to generate token")
		return nil
	}

	return &models.LoginResponse{
		ID:       int(user.ID),
		Email:    user.Email,
		UserName: user.UserName,
		Token: models.TokenStruct{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (a *AuthService) Verify(c *gin.Context) bool {
	var reqBody *models.UserVerifyOtp

	// get and validate
	if err := c.ShouldBindBodyWithJSON(&reqBody); err != nil {
		response.BadRequestErrorWithFields(c, response.ErrCodeValidation, utils.GetObjMessage(err))
		return false
	}

	//  check user exists
	user, err := repo.GetDetailUserByEmail(global.MDB, reqBody.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			response.BadRequestError(c, response.ErrCodeResourceConflict, err.Error())
			return false
		} else {
			response.BadRequestError(c, response.ErrCodeResourceConflict, "Email not found")
			return false
		}
	}

	// check user is active
	if user.IsActive {
		response.ForbiddenError(c, response.ErrCodeForbidden, "User is already active")
		return false
	}

	// check otp
	otp, err := redis.GetOtpKey(c, global.RDB, reqBody.Email)
	if err != nil {
		response.BadRequestError(c, response.ErrCodeResourceConflict, err.Error())
		return false
	}
	if otp == "" {
		response.BadRequestError(c, response.ErrCodeResourceConflict, "OTP is expired, please resend OTP")
		return false
	}

	// compare otp
	if otp != fmt.Sprintf("%d", reqBody.Otp) {
		response.BadRequestError(c, response.ErrCodeResourceConflict, "OTP is incorrect")
		return false
	}

	if err := repo.ActiveUser(global.MDB, user); err != nil {
		logger.LogError(fmt.Sprintf("Failed to active user %s: %v\n", user.Email, err))
		response.InternalServerError(c, response.ErrCodeInternalServer, "Failed to active user")
		return false
	}

	// clear otp of this email from cache
	go func() {
		if err := redis.DeleteOtpKey(c, global.RDB, reqBody.Email); err != nil {
			logger.LogError(fmt.Sprintf("Failed to delete OTP of user %s: %v\n", user.Email, err))
		}
	}()

	return true
}

func (a *AuthService) ResendOtp(c *gin.Context) bool {
	var reqBody *models.UserResendOtp

	// get and validate
	if err := c.ShouldBindBodyWithJSON(&reqBody); err != nil {
		response.BadRequestErrorWithFields(c, response.ErrCodeValidation, utils.GetObjMessage(err))
		return false
	}

	// get user by email
	user, err := repo.GetDetailUserByEmail(global.MDB, reqBody.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			response.BadRequestError(c, response.ErrCodeDBConnection, err.Error())
			return false
		} else {
			response.BadRequestError(c, response.ErrCodeResourceConflict, "Email not found")
			return false
		}
	}

	// check user is active
	if user.IsActive {
		response.ForbiddenError(c, response.ErrCodeForbidden, "User is already active")
		return false
	}

	// check otp already sent
	if err := checkOtpAlreadySent(c, reqBody.Email); err != nil {
		response.BadRequestError(c, response.ErrCodeResourceConflict, err.Error())
		return false
	}

	// generate otp
	otp := helpers.GenerateOTP(6)

	// send email verification
	otpMailData := models.DataOtpMail{
		Title:  constants.TitleOtpMail,
		To:     reqBody.Email,
		Expire: helpers.FormatDuration(constants.ExpiresOTP),
		Otp:    otp,
		Name:   user.UserName,
	}

	// save otp to cache
	if err := redis.SaveOtpKey(c, global.RDB, reqBody.Email, otp); err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer, "Failed to save OTP")
		return false
	}

	if err := mailer.SendOptMail(otpMailData); err != nil {
		response.InternalServerError(c, response.ErrCodeInternalServer, "Failed to send email verification")
		// clear otp of this email from cache if send error
		redis.DeleteOtpKey(c, global.RDB, reqBody.Email)
		return false
	}

	return true
}

func checkOtpAlreadySent(c *gin.Context, email string) error {
	otp, err := redis.GetOtpKey(c, global.RDB, email)
	if err != nil {
		return err
	}

	if otp != "" {
		return errors.New("OTP is already sent")
	}

	return nil
}
