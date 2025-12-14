package handler

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/AKASHPRADHAN8336/aniyxProject/internal/models"
	"github.com/AKASHPRADHAN8336/aniyxProject/internal/service"
)

type UserHandler struct {
	svc      service.UserService
	logger   *zap.Logger
	validate *validator.Validate
}

func NewUserHandler(svc service.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		svc:      svc,
		logger:   logger,
		validate: validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := h.svc.CreateUser(c.Context(), req)
	if err != nil {
		h.logger.Error("Failed to create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Invalid user ID", zap.String("id", idStr), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := h.svc.GetUser(c.Context(), int32(id))
	if err != nil {
		h.logger.Error("User not found", zap.Int32("id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	h.logger.Info("ListUsers handler called")

	users, err := h.svc.ListUsers(c.Context())
	if err != nil {
		h.logger.Error("Failed to list users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to list users",
		})
	}

	h.logger.Info("Returning users", zap.Int("count", len(users)))

	if users == nil {
		users = []*models.UserResponse{}
	}

	return c.JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Invalid user ID", zap.String("id", idStr), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := h.svc.UpdateUser(c.Context(), int32(id), req)
	if err != nil {
		h.logger.Error("Failed to update user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Invalid user ID", zap.String("id", idStr), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	err = h.svc.DeleteUser(c.Context(), int32(id))
	if err != nil {
		h.logger.Error("Failed to delete user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
