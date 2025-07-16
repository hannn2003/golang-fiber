package controllers

import (
	"fiber-demo/models"
	"fiber-demo/queries"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	Repo *queries.UserRepository
}

func NewUserController(repo *queries.UserRepository) *UserController {
	return &UserController{Repo: repo}
}


// GetUserById retrieves a user by their ID
// @Summary Get user by ID
// @Description Retrieve a user by their unique ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /api/v1/users/{id} [get]
func (ctl *UserController) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	idParse, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}

	user, err := ctl.Repo.GetUserById(idParse)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// GetUsers retrieves all users
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /api/v1/users [get]
func (ctl *UserController) GetUsers(c *fiber.Ctx) error {
	users, err := ctl.Repo.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}
	return c.JSON(users)
}


// UserSignUp handles user registration
// It expects a JSON body with username, password, and email fields.
// If successful, it returns the created user with a 201 status code.
// @Summary User Sign Up
// @Description Register a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 201 {object} models.User
// @Router /api/v1/create-user [post]
func (ctl *UserController) UserSignUp(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if input.Username == "" || input.Password == "" || input.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username, password, and email are required"})
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := ctl.Repo.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}