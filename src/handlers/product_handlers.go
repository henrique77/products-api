package handlers

import (
	"net/http"
	"strconv"

	"github.com/henrique77/product-api/src/modules/product/entities"
	service "github.com/henrique77/product-api/src/modules/product/services"
	"github.com/henrique77/product-api/src/utils"
	"github.com/labstack/echo/v4"
)

type (
	ProductHandler interface {
		GetProducts(c echo.Context) error
		CreateProduct(c echo.Context) error
		GetProduct(c echo.Context) error
		UpdateProduct(c echo.Context) error
		DeleteProduct(c echo.Context) error
	}

	productHandler struct {
		service.ProductService
	}
)

func (h *productHandler) GetProducts(c echo.Context) error {
	r, err := h.ProductService.GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

func (h *productHandler) CreateProduct(c echo.Context) error {
	var p *entities.Products

	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}

	err := h.ProductService.CreateProduct(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "Produto criado!")
}

func (h *productHandler) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	r, err := h.ProductService.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

func (h *productHandler) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var p *entities.Products

	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "args is invalid"})
	}

	r, err := h.ProductService.UpdateProduct(id, p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

func (h *productHandler) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.ProductService.DeleteProduct(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "OK")
}
