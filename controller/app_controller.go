package controller

import (
	"github.com/dikyayodihamzah/simrs/handler"
	"github.com/gofiber/fiber/v2"
)

type AppController interface {
	Route(app *fiber.App)
}

type appController struct {
	AppHandler handler.AppHandler
}

func NewAppController(appHandler handler.AppHandler) AppController {
	return &appController{
		AppHandler: appHandler,
	}
}

func (controller *appController) Route(app *fiber.App) {
	app.Static("", "./static/assets")

	app.Get("/login", handler.Login)
	app.Get("/", controller.AppHandler.Dashboard)
	app.Get("/data-pasien", func(c *fiber.Ctx) error {
		return c.Render("pages-data-pasien", fiber.Map{})
	})
}
