package main

import (
	"os"
	"simple-golang-social-media-app/internal/handler"
	"simple-golang-social-media-app/internal/middleware"
	"simple-golang-social-media-app/internal/model"
	"simple-golang-social-media-app/internal/repository"
	"simple-golang-social-media-app/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&model.User{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*") // Add this line to enable HTML template rendering

	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/home", func(c *gin.Context) {
			email, _ := c.Get("email")
			user, err := userService.FindByEmail(email.(string))
			if err != nil {
				c.JSON(404, gin.H{"error": "User not found"})
				return
			}
			c.JSON(200, gin.H{"message": "Hello " + user.Username})
		})
		protected.POST("/logout", userHandler.Logout)
	}

	r.Run(":8080")
}
