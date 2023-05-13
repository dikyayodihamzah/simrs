package controller

import (
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

func Route(app *fiber.App) {
	app.Static("/", "./static/assets")

	app.Get("/", Home)
	app.Get("/data-pasien", func(c *fiber.Ctx) error {
		return c.Render("pages-data-pasien", fiber.Map{})
	})
}

func Home(c *fiber.Ctx) error {
	if err := c.Render("home", fiber.Map{}); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "can't render html")
	}

	return nil
}
