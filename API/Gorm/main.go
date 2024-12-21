package gorm

import (
	"giuthub.com./gofiber/fiber"
)



func helloWord(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func mian() {
	app := fiber.New()

	app.Get("/", helloWord)

	app.Listen(":3000")


}