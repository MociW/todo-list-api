package main

import (
	"todo-list-api/internal/controller"
	"todo-list-api/internal/database"
	"todo-list-api/internal/repository"
	"todo-list-api/internal/server"
	"todo-list-api/internal/service"
)

func main() {
	db := database.NewDB()
	userRepository := repository.NewUserRepository()
	addressRepository := repository.NewAddressRepository()
	todoRepository := repository.NewTodoRepository()

	userService := service.NewUserService(db, userRepository)
	addressService := service.NewAddressService(db, addressRepository, userRepository)
	todoService := service.NewTodoService(db, todoRepository, userRepository)

	userController := controller.NewUserController(userService)
	addressController := controller.NewAddressController(addressService)
	todoController := controller.NewTodoController(todoService)

	routeConfig := server.RouteConfig{
		UserController:    userController,
		AddressController: addressController,
		TodoController:    todoController,
	}

	app := routeConfig.NewRouter()
	app.Listen(":3000")
}
