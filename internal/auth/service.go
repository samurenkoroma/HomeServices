package auth

import (
	"errors"
	"samurenkoroma/services/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *user.UserRepository
}

func NewAuthService(userRepo *user.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepo,
	}
}

func (service *AuthService) updateRefreshToken(email, refresh string) error {
	existedUser, _ := service.userRepository.FindByEmail(email)
	return service.userRepository.Update(existedUser.Email, &user.User{RefreshToken: refresh})
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.userRepository.FindByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &user.User{
		Email:    email,
		PassHash: string(hashedPassword),
		Name:     name,
	}

	_, err = service.userRepository.Create(user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}

func (service *AuthService) Login(email, pass string) (string, error) {
	existedUser, _ := service.userRepository.FindByEmail(email)
	if existedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existedUser.PassHash), []byte(pass)); err != nil {
		return "", errors.New(ErrWrongCredentials)
	}
	return existedUser.Email, nil
}

func (service *AuthService) Refresh(refresh string) error {
	existedUser, _ := service.userRepository.FindByRefresh(refresh)
	if existedUser == nil {
		return errors.New(ErrWrongCredentials)
	}
	return service.userRepository.Update(existedUser.Email, &user.User{RefreshToken: refresh})
}
