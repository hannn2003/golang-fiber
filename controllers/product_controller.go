package controllers

import (
	"fiber-demo/models"

	"github.com/gofiber/fiber/v2"
)

// GetProducts returns a list of products
// @Summary Get list of products
// @Description Returns a list of products available in the store
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /api/v1/products [get]
func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{
		{Name: "Product 1"},
		{Name: "Product 2"},
		{Name: "Product 3"},
	}
	return c.JSON(products)
}