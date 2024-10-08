package constants

import "time"

const (
	SpamKey           = "spam_user"
	SpamKeyLogin      = "spam_user_login"
	SpamKeyOTP        = "spam_user_otp"
	OTPKey            = "otp_user"
	AccessTokenBlack  = "access_token_black"
	RefreshTokenBlack = "refresh_token_black"
)

const (
	RequestThreshold = 5
)

const (
	InitialBlock   = 5 * time.Minute
	ExtendBlock    = 30 * time.Minute
	ExpireDuration = 30 * time.Second
	ExpireSevenDay = 7 * 24 * time.Hour
)
