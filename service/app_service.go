package service

import "github.com/gofiber/fiber/v2"

// type AppService interface {

// }

// type appService struct {

// }

func Dashboard(c *fiber.Ctx) error {
	return c.Render("pages-dashboard", fiber.Map{
		"Title":        "Dashboard",
		"UsedRoom":     "145",
		"TotalRoom":    "145",
		"TotalPatient": "145",
	})
}

func Login(c *fiber.Ctx) error {
	return c.Render("pages-login", fiber.Map{})
}

func InformasiKamar(c *fiber.Ctx) error {
	return c.Render("pages-data-pasien", fiber.Map{})
}

func DataPasien(c *fiber.Ctx) error {
	return c.Render("dashboard", fiber.Map{})
}

func About(c *fiber.Ctx) error {
	return c.Render("dashboard", fiber.Map{})
}

func Kontak(c *fiber.Ctx) error {
	return c.Render("pages-contact", fiber.Map{})
}
