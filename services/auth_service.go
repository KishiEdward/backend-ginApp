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