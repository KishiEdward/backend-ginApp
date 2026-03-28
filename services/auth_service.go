package services

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/KishiEdward/backend_ginApp/config"
	"github.com/KishiEdward/backend_ginApp/models"
	"github.com/KishiEdward/backend_ginApp/repositories"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{userRepo: repositories.NewUserRepository()}
}

func (s *AuthService) generateJWT(user *models.User) (string, error) {
	expireHours, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOURS"))
	if expireHours == 0 {
		expireHours = 24
	}

	claims := jwt.MapClaims{
		"sub":            user.ID,
		"firebase_uid":   user.FirebaseUID,
		"email":          user.Email,
		"name":           user.Name,
		"role":           user.Role,
		"iat":            time.Now().Unix(),
		"exp":            time.Now().Add(time.Hour * time.Duration(expireHours)).Unix(),
		"email_verified": user.EmailVerified,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}