package render

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	if err := c.Render("home", fiber.Map{}); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "can't render html")
	}

	return nil
}
