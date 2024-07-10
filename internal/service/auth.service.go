package service

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Register() {
	// do something
}
