package middlewares

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup(app *fiber.App) {
	// panic recovery
	app.Use(recover.New())

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Change to specific origins if needed
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	// limit repeated requests
	maxRequest := 50
	app.Use(Limit(maxRequest, 2))

	// http request logger
	app.Use(logger.New(logger.Config{
		Format:     `${time} ${locals:requestid} ${status} - ${method} ${url}` + "\n\n",
		TimeFormat: "2006/01/02 15:04:05",
	}))

	// debugger
	debug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if debug {
		app.Use(pprof.New())
	}
}

func Limit(maxRequest int, duration time.Duration) func(*fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        maxRequest,
		Expiration: duration * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(fiber.Map{
				"error":   true,
				"message": "Too many requests",
			})
		},
	})
}

func MaxBodySize(sizeInMB int) fiber.Handler {
	sizeInMB = sizeInMB * 1024 * 1024
	return func(c *fiber.Ctx) error {
		if len(c.Body()) >= sizeInMB {
			// custom response here
			return fiber.ErrRequestEntityTooLarge
		}
		return c.Next()
	}
}
