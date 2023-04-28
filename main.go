package main

import (
	"fmt"
	"time"

	"github.com/dikyayodihamzah/simrs.git/config"
	"github.com/dikyayodihamzah/simrs.git/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	time.Local = time.UTC

	config.ConnectDB()
	fmt.Println("Connected to database")

	serverConfig := config.NewServerConfig()

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	err := app.Listen(serverConfig.URI)
	exception.PanicIfError(err)
}
