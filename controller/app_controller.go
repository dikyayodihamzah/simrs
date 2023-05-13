package controller

import (
	"github.com/dikyayodihamzah/simrs.git/render"
	"github.com/gofiber/fiber/v2"
)

// type AppController interface {
// 	NewAppRouter(app *fiber.App)
// }

// type appController struct {
// 	AppService service.AppService
// }

// func NewAppController(appService service.AppService) AppController {
// 	return &appController{
// 		AppService: appService,
// 	}
// }

func Controller(app *fiber.App) {
	app.Static("/", "./static/assets")

	app.Get("/", render.Home)

	app.Get("/data-pasien", func(c *fiber.Ctx) error {
		return c.Render("pages-data-pasien", fiber.Map{})
	})
}
