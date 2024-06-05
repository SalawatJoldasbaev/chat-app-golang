package main

import (
	"github.com/SalawatJoldasbaev/chat-app-golang/api"
	"github.com/SalawatJoldasbaev/chat-app-golang/api/middlewares"
	"github.com/SalawatJoldasbaev/chat-app-golang/configs"
	"github.com/SalawatJoldasbaev/chat-app-golang/internal/database"
	"github.com/SalawatJoldasbaev/chat-app-golang/pkg/utility"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	config := configs.Load(".")
	utility.ZapLogger(config.App.Env)
	//setup database
	database.ConnectDatabase(config)
	//fiber config
	app := fiberConfig(config)
	//setup middleware
	middlewares.Setup(app)
	//setup api routes
	api.Setup(app, database.DB)
	// listen app
	if err := app.Listen(config.App.Host + ":" + config.App.Port); err != nil {
		utility.Logger.Fatal("Error while listening to port", zap.Error(err))
	}
}

func fiberConfig(config *configs.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       config.App.Name,
	})
	return app
}
