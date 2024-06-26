package main

import (
	"goHex/internal/core/services"
	"goHex/internal/handlers"
	"goHex/internal/repositories"
	"goHex/internal/server"
)

func main() {
	mongoConn := "secret"
	//repositories
	userRepository := repositories.NewUserRepository(mongoConn)
	//services
	userService := services.NewUserService(userRepository)
	//handlers
	userHandlers := handlers.NewUserHandlers(userService)
	//server
	httpServer := server.NewServer(
		userHandlers,
	)
	httpServer.Initialize()
}
