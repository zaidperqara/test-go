package main

import (
	"github.com/aenmurtic/be-hijooin-admin/internal/auth"
	"github.com/joho/godotenv"
	"os"

	"github.com/aenmurtic/be-hijooin-admin/internal/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load() // Load .env file at the project root
	if err != nil {
		panic("Failed to load .env file")
	}

	// Database setup
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto-migrate (if needed)
	err = db.AutoMigrate(&user.User{})
	if err != nil {
		return
	}

	// Initialize repository
	userRepo := user.NewUserRepository(db)

	// Initialize service
	userService := user.NewUserService(userRepo)

	// Initialize handler
	userHandler := user.NewUserHandler(userService)

	// Initialize Gin router
	router := gin.Default()

	// Register route
	router.POST("/register", userHandler.Register)

	// Register login route
	router.POST("/login", userHandler.Login)

	// Protected routes example
	protectedRoutes := router.Group("/admin")
	protectedRoutes.Use(auth.AuthorizeJWT())
	//protectedRoutes.GET("/dashboard", someAdminHandler)

	err = router.Run()
	if err != nil {
		return
	} // Start the server (modify port if needed)
}
