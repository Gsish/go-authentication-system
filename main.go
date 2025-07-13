package main

import (
	"jwt-auth-service/handlers"
	"jwt-auth-service/middleware"
	"jwt-auth-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB connection - change URI if needed
	models.InitMongo("mong url")

	r := gin.Default()

	r.POST("/register", handlers.RegisterHandler)
	r.POST("/login", handlers.LoginHandler)

	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/protected", handlers.ProtectedHandler)
	}

	r.Run(":8080")
}
