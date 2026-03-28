package repositories

import (
	"github.com/KishiEdward/backend_ginApp/config"
	"github.com/KishiEdward/backend_ginApp/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}