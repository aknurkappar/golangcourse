package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/api"
	"main/middleware"
)

func main() {
	r := gin.Default()
	api.Connect()
	api.InitRoles()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	r.Use(cors.New(config))

	r.POST("/register", api.RegisterUser)
	r.POST("/login", api.Login)
	r.POST("/users", middleware.JWTMiddleware(middleware.AdminMiddleware(api.CreateUser)))
	r.GET("/users", middleware.JWTMiddleware(middleware.AdminMiddleware(api.GetUsers)))
	r.PUT("/users", middleware.JWTMiddleware(middleware.AdminMiddleware(api.UpdateUser)))
	r.GET("/roles", middleware.JWTMiddleware(middleware.AdminMiddleware(api.GetRoles)))
	r.GET("/profile", middleware.JWTMiddleware(api.GetUserProfile))
	r.Run() // 8080
}
