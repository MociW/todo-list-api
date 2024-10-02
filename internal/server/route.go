package server

import (
	"todo-list-api/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	UserController    *controller.UserController
	AddressController *controller.AddressController
	TodoController    *controller.TodoController
}

func (c *RouteConfig) NewRouter() *fiber.App {
	app := fiber.New()
	//users
	app.Post("/v1/users", c.UserController.Register)
	app.Post("/v1/users/_login", c.UserController.Login)
	app.Post("/v1/users/_logout", c.UserController.Logout)
	app.Put("/v1/users/_current", c.UserController.Update)
	app.Delete("/v1/users/_current", c.UserController.Delete)

	//addresses
	app.Get("/v1/users/addresses", c.AddressController.List)
	app.Get("/v1/users/addresses/:addressId", c.AddressController.Get)
	app.Post("/v1/users/addresses", c.AddressController.Create)
	app.Put("/v1/users/addresses/:addressId", c.AddressController.Update)
	app.Delete("/v1/users/addresses/:addressId", c.AddressController.Delete)

	//todos
	app.Get("/v1/users/todos", c.TodoController.List)
	app.Get("/v1/users/todos/:todoId", c.TodoController.Get)
	app.Post("/v1/users/todos", c.TodoController.Create)
	app.Put("/v1/users/todos/:todoId", c.TodoController.Update)
	app.Put("/v1/users/todos/status/:todoId", c.TodoController.UpdateStatus)
	app.Delete("/v1/users/todos/:todoId", c.TodoController.Delete)

	return app
}
