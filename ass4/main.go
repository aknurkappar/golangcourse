package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"main/api"
	"main/middleware"
	"net/http"
)

func testHandler(c *gin.Context) {
	c.Header("X-Custom-Header", "TestValue")
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

var Logger *zap.Logger

func initLogger() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger")
	}
}

func main() {
	r := gin.Default()

	secretKey := []byte("32-byte-long-secret-key")
	csrfMiddleware := csrf.Protect(secretKey, csrf.Secure(false))
	r.Use(func(c *gin.Context) {
		csrfMiddleware(http.DefaultServeMux)
		c.Next()
	})
	r.Use(middleware.SecurityHeadersMiddleware())

	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "status"},
	)
	prometheus.MustRegister(requestCounter)
	r.Use(func(c *gin.Context) {
		c.Next()
		requestCounter.WithLabelValues(c.Request.Method, http.StatusText(c.Writer.Status())).Inc()
	})

	initLogger()
	defer Logger.Sync()
	r.Use(middleware.LogUserActionMiddleware(Logger))

	api.Connect()
	api.InitRoles()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	r.Use(cors.New(config))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.POST("/register", api.RegisterUser)
	r.POST("/login", api.Login)
	r.POST("/users", middleware.JWTMiddleware(middleware.AdminMiddleware(api.CreateUser)))
	r.GET("/users", middleware.JWTMiddleware(middleware.AdminMiddleware(api.GetUsers)))
	r.PUT("/users", middleware.JWTMiddleware(middleware.AdminMiddleware(api.UpdateUser)))
	r.GET("/roles", middleware.JWTMiddleware(middleware.AdminMiddleware(api.GetRoles)))
	r.GET("/profile", middleware.JWTMiddleware(api.GetUserProfile))
	r.GET("/test", testHandler)
	r.Run() // 8080
}
