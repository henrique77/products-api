package service

import (
	"github.com/henrique77/product-api/src/modules/product/entities"
	"github.com/henrique77/product-api/src/modules/product/repositories"
)

type (
	ProductService interface {
		GetProducts() ([]*entities.Products, error)
		CreateProduct(*entities.Products) error
		GetProduct(int) (*entities.Products, error)
		UpdateProduct(int, *entities.Products) (*entities.Products, error)
		DeleteProduct(int) error
	}

	productService struct {
		repository *repositories.Repository
	}
)

func (s *productService) GetProducts() ([]*entities.Products, error) {
	r, err := s.repository.Product.Get()
	return r, err
}

func (s *productService) CreateProduct(p *entities.Products) error {
	err := s.repository.Product.Create(p)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) GetProduct(id int) (*entities.Products, error) {
	r, err := s.repository.Product.GetOne(id)
	return r, err
}

func (s *productService) UpdateProduct(id int, p *entities.Products) (*entities.Products, error) {
	r, err := s.repository.Product.Update(id, p)
	return r, err
}

func (s *productService) DeleteProduct(id int) error {
	err := s.repository.Product.Delete(id)
	return err
}
