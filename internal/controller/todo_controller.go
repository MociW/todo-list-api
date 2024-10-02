package controller

import (
	"strconv"
	"todo-list-api/internal/config"
	"todo-list-api/internal/model"
	"todo-list-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TodoController struct {
	TodoService *service.TodoService
}

func NewTodoController(todoService *service.TodoService) *TodoController {
	return &TodoController{todoService}
}

func (t *TodoController) Create(c *fiber.Ctx) error {
	request := new(model.CreateTodoRequest)
	err := c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Salt), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse claims",
		})
	}

	if (*claims)["id"].(string) != request.UserId {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	response, err := t.TodoService.Create(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}

func (t *TodoController) Update(c *fiber.Ctx) error {
	request := new(model.UpdateTodoRequest)
	todoId, err := strconv.Atoi(c.Params("todoId"))
	if err != nil {
		return err
	}
	request.ID = uint(todoId)

	err = c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Salt), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid token",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Token validation failed",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "InternalServerError: Failed to parse token claims",
		})
	}

	jwtId, ok := (*claims)["id"].(string)
	if !ok || jwtId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "BadRequest: User ID not found in token",
		})
	}

	if jwtId != request.UserId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden: You do not have access to this data",
		})
	}

	response, err := t.TodoService.Update(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}

func (t *TodoController) UpdateStatus(c *fiber.Ctx) error {
	request := new(model.UpdateStatusTodoRequest)
	todoId, err := strconv.Atoi(c.Params("todoId"))
	if err != nil {
		return err
	}
	request.ID = uint(todoId)

	err = c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Salt), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid token",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Token validation failed",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "InternalServerError: Failed to parse token claims",
		})
	}

	jwtId, ok := (*claims)["id"].(string)
	if !ok || jwtId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "BadRequest: User ID not found in token",
		})
	}

	if jwtId != request.UserId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden: You do not have access to this data",
		})
	}

	response, err := t.TodoService.UpdateStatus(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}

func (t *TodoController) Delete(c *fiber.Ctx) error {
	request := new(model.DeleteTodoRequest)
	todoId, err := strconv.Atoi(c.Params("todoId"))
	if err != nil {
		return err
	}
	request.ID = uint(todoId)

	err = c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Salt), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid token",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Token validation failed",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "InternalServerError: Failed to parse token claims",
		})
	}

	jwtId, ok := (*claims)["id"].(string)
	if !ok || jwtId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "BadRequest: User ID not found in token",
		})
	}

	if jwtId != request.UserId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden: You do not have access to this data",
		})
	}

	err = t.TodoService.Delete(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Delete User Successfully",
	})
}
func (t *TodoController) Get(c *fiber.Ctx) error {
	request := new(model.GetTodoRequest)
	todoId, err := strconv.Atoi(c.Params("todoId"))
	if err != nil {
		return err
	}
	request.ID = uint(todoId)

	err = c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Salt), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid token",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Token validation failed",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "InternalServerError: Failed to parse token claims",
		})
	}

	jwtId, ok := (*claims)["id"].(string)
	if !ok || jwtId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "BadRequest: User ID not found in token",
		})
	}

	if jwtId != request.UserId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden: You do not have access to this data",
		})
	}

	response, err := t.TodoService.Get(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}
func (t *TodoController) List(c *fiber.Ctx) error {
	request := new(model.ListTodoRequest)
	err := c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Salt), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid token",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Token validation failed",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "InternalServerError: Failed to parse token claims",
		})
	}

	jwtId, ok := (*claims)["id"].(string)
	if !ok || jwtId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "BadRequest: User ID not found in token",
		})
	}

	if jwtId != request.UserId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden: You do not have access to this data",
		})
	}

	response, err := t.TodoService.List(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}
