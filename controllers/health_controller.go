package controllers

import "github.com/gofiber/fiber/v2"

// HealthApi returns the health status of the API
// @Summary Get API health status
// @Description Returns the health status of the API
// @Tags Health
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/ [get]
func HealthApi(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "200",
		"message": "API is running smoothly",
	})
}