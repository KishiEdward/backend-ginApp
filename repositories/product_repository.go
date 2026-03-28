package repositories 

import (
	"github.com/KishiEdward/backend_ginApp/config"
	"github.com/KishiEdward/backend_ginApp/models"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}
