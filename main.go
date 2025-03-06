package main

import (
	"fiber-generic/handlers"
	"fiber-generic/models"
	"fiber-generic/repositories"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=gin_generic port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.User{}) // Auto migrate model User

	// Inisialisasi repository dan handler
	userRepo := &repositories.GenericRepository[models.User]{DB: db}
	userHandler := &handlers.UserHandler{Repo: userRepo}

	// Inisialisasi Go-Fiber
	app := fiber.New()

	// Routing
	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Post("/users", userHandler.CreateUser)

	// Jalankan server
	app.Listen(":8080")
}
