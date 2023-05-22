package main

import (
	"log"
	"time"

	"github.com/dikyayodihamzah/simrs/config"
	"github.com/dikyayodihamzah/simrs/controller"
	"github.com/dikyayodihamzah/simrs/handler"
	"github.com/dikyayodihamzah/simrs/repository"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	time.Local = time.UTC

	engine := html.New("./static", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "base-layout",
	})

	serverConfig := config.NewServerConfig()
	db := config.NewDB()
	validate := validator.New()

	roomRespoitory := repository.NewRoomRepository(db)
	impatientRespoitory := repository.NewImpatientRepository(db)
	appHandler := handler.NewAppHandler(roomRespoitory, impatientRespoitory, validate)
	appController := controller.NewAppController(appHandler)

	appController.Route(app)

	err := app.Listen(serverConfig.URI)
	log.Fatal(err)
}
