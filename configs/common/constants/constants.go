package constants

const (
	ExpiresAccessToken  = 3 * 60 * 60      // 3 hours
	ExpiresRefreshToken = 7 * 24 * 60 * 60 // 7 days
	ExpiresOTP          = 5 * 60           // 5 minutes
)

const (
	AuthorizationHeader = "Authorization"
	RefreshTokenHeader  = "RefreshToken"
)

const (
	DevMode  = "development"
	ProdMode = "production"
)

const TitleOtpMail = "OTP Verification"
