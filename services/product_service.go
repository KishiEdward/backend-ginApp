package services

import (
	"github.com/KishiEdward/backend_ginApp/models"
	"github.com/KishiEdward/backend_ginApp/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{productRepo: repositories.NewProductRepository()}
}