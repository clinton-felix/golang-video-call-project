package handlers

import "github.com/gofiber/fiber/v2"

// welcome function runs when the root route is hit
// it renders the layout/main in the views directory
func Welcome(c *fiber.Ctx) error {
	return c.Render("welcome", nil , "layouts/main")
}