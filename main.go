package main

import (
	"fmt"
	"time"

	"github.com/dikyayodihamzah/simrs.git/config"
	"github.com/dikyayodihamzah/simrs.git/controller"
	"github.com/dikyayodihamzah/simrs.git/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	time.Local = time.UTC

	config.ConnectDB()
	fmt.Println("Connected to database")

	serverConfig := config.NewServerConfig()

	engine := html.New("./static", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		ViewsLayout: "base-layout",
	})

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	controller.Route(app)

	err := app.Listen(serverConfig.URI)
	exception.PanicIfError(err)
}
