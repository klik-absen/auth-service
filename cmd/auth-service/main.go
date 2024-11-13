package main

import (
	"ka-auth-service/internal/application/service"
	"ka-auth-service/internal/infrastructure/db"
	"ka-auth-service/internal/infrastructure/env"
	"ka-auth-service/internal/infrastructure/repository"
	"ka-auth-service/internal/interfaces/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config := env.LoadConfig()
	dbConn := db.NewDBConnection(config)

	userRepo := repository.NewUserRepoImpl(dbConn)
	userSessionRepo := repository.NewUserSessionRepoImpl(dbConn)
	authService := service.NewAuthService(userRepo, userSessionRepo)
	authController := controller.NewAuthController(authService)

	router := gin.Default()
	router.POST("/api/v1/auth", authController.Authentication)

	log.Println("Server is running on port 3000")
	log.Fatal(router.Run(":3000"))
}
