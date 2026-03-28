package repositories

import (
	"github.com/KishiEdward/backend_ginApp/config"
	"github.com/KishiEdward/backend_ginApp/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByFirebaseUID(uid string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("firebase_uid = ?", uid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}