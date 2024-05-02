package service

import (
	"github.com/henrique77/product-api/src/modules/product/repositories"
)

type Services struct {
	Product ProductService
}

func New(s *repositories.Repository) *Services {
	return &Services{
		Product: &productService{repository: s},
	}
}
