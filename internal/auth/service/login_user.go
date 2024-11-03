package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"shortener-smile/internal/auth/repository"
	"shortener-smile/internal/common"
)

type LoginUserService struct {
	repo       repository.UserRepository
	appCtx     *common.ApplicationContext
	jwtService *JWTService
}

func NewLoginUserService(
	repo repository.UserRepository,
	ctx *common.ApplicationContext,
	jwtService *JWTService,
) *LoginUserService {
	return &LoginUserService{
		repo:       repo,
		appCtx:     ctx,
		jwtService: jwtService,
	}
}

func (s *LoginUserService) CreateJwtByLoginAndPassword(login, password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	fmt.Println(login, string(hashedPass))

	user, err := s.repo.GetUserByLogin(login)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	return s.jwtService.CreateJwtForId(user.Id)
}
