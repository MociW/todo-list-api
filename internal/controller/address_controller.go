package controller

import (
	"strconv"
	"todo-list-api/internal/config"
	"todo-list-api/internal/model"
	"todo-list-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AddressController struct {
	AddressService *service.AddressService
}

func NewAddressController(addressService *service.AddressService) *AddressController {
	return &AddressController{AddressService: addressService}
}

func (a *AddressController) Create(c *fiber.Ctx) error {
	request := new(model.CreateAddressRequest)
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

	response, err := a.AddressService.Create(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}

func (a *AddressController) Update(c *fiber.Ctx) error {
	request := new(model.UpdateAddressRequest)
	addressId, err := strconv.Atoi(c.Params("addressId"))
	if err != nil {
		return err
	}
	request.ID = uint(addressId)

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

	response, err := a.AddressService.Update(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}

func (a *AddressController) Delete(c *fiber.Ctx) error {
	request := new(model.DeleteAddressRequest)
	addressId, err := strconv.Atoi(c.Params("addressId"))
	if err != nil {
		return err
	}
	request.ID = uint(addressId)

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

	err = a.AddressService.Delete(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Delete User Successfully",
	})
}

func (a *AddressController) Get(c *fiber.Ctx) error {
	request := new(model.GetAddressRequest)
	addressId, err := strconv.Atoi(c.Params("addressId"))
	if err != nil {
		return err
	}
	request.ID = uint(addressId)

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

	response, err := a.AddressService.Get(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}

func (a *AddressController) List(c *fiber.Ctx) error {
	request := new(model.ListAddressRequest)
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

	response, err := a.AddressService.List(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}
