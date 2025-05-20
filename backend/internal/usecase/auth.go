package usecase

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthUseCase interface {
	Register(user *entity.User) error
	Login(credentials *entity.AuthRequest) (*entity.AuthResponse, error)
}

type authUseCase struct {
	userRepo repository.UserRepository
}

func NewAuthUseCase(userRepo repository.UserRepository) AuthUseCase {
	return &authUseCase{userRepo: userRepo}
}

func (uc *authUseCase) Register(user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return uc.userRepo.Create(user)
}

func (uc *authUseCase) Login(credentials *entity.AuthRequest) (*entity.AuthResponse, error) {
	user, err := uc.userRepo.FindByLogin(credentials.Login)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
	//	return nil, fmt.Errorf("invalid password: %v", err)
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &entity.AuthResponse{Token: tokenString}, nil
}
