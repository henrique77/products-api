package handlers

import (
	service "github.com/henrique77/product-api/src/modules/product/services"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	ProductHandler
}

func New(s *service.Services) *Handlers {
	return &Handlers{
		ProductHandler: &productHandler{s.Product},
	}
}

func SetApi(e *echo.Echo, h *Handlers) {
	g := e.Group("/api/v1")

	g.GET("/product", h.ProductHandler.GetProducts)
	g.POST("/product", h.ProductHandler.CreateProduct)
	g.GET("/product/:id", h.ProductHandler.GetProduct)
	g.PUT("/product/:id", h.ProductHandler.UpdateProduct)
	g.DELETE("/product/:id", h.ProductHandler.DeleteProduct)
}
