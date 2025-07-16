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

// PostProduct creates a new product
// @Summary Create a new product
// @Description Creates a new product with the provided details
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product details"
// @Router /api/v1/products [post]
func PostProduct(c* fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser((&product)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": product,
	})
}

