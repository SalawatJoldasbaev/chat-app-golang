package api

import (
	"github.com/SalawatJoldasbaev/chat-app-golang/api/middlewares"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/app/handlers"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/app/services"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api", middlewares.JwtMiddleware())
	//repositories
	userRepo := repositories.NewUserRepository(db)

	//handlers
	userHandler := handlers.NewUserHandler(services.NewUserService(userRepo))
	//groups
	auth := api.Group("/auth")
	users := api.Group("/users")

	//auth
	auth.Post("/register", userHandler.Register)
	auth.Post("/login", userHandler.Login)
	auth.Get("/me", userHandler.GetMe)
	auth.Post("/logout", userHandler.Logout)

	//users
	users.Patch("/", userHandler.Update)
	users.Put("/change-password", userHandler.ChangePassword)
	users.Delete("/", userHandler.Delete)
}
