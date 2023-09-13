package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func DefaulHandler(c *fiber.Ctx) error {
	path := c.Path()
	return c.SendString(path)
}

func InitServer() {
	app := fiber.New()
	// Initialize default config
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\nRequest Body:\n${body}\n\nResponse:\n${resBody}\n\n",
	}))
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// User Routes
	user := v1.Group("/user")
	user.Get("/", DefaulHandler)
	user.Post("/", DefaulHandler)
	user.Patch("/", DefaulHandler)
	user.Delete("/", DefaulHandler)

	attachment := v1.Group("/attachment")
	attachment.Get("/", DefaulHandler)
	attachment.Post("/", DefaulHandler)
	attachment.Patch("/", DefaulHandler)
	attachment.Delete("/", DefaulHandler)

	board := v1.Group("/board")
	board.Get("/", DefaulHandler)
	board.Post("/", DefaulHandler)
	board.Patch("/", DefaulHandler)
	board.Delete("/", DefaulHandler)

	card_label := v1.Group("/card_label")
	card_label.Get("/", DefaulHandler)
	card_label.Post("/", DefaulHandler)
	card_label.Patch("/", DefaulHandler)
	card_label.Delete("/", DefaulHandler)

	card := v1.Group("/card")
	card.Get("/", DefaulHandler)
	card.Post("/", DefaulHandler)
	card.Patch("/", DefaulHandler)
	card.Delete("/", DefaulHandler)

	comment := v1.Group("/comment")
	comment.Get("/", DefaulHandler)
	comment.Post("/", DefaulHandler)
	comment.Patch("/", DefaulHandler)
	comment.Delete("/", DefaulHandler)

	label := v1.Group("/label")
	label.Get("/", DefaulHandler)
	label.Post("/", DefaulHandler)
	label.Patch("/", DefaulHandler)
	label.Delete("/", DefaulHandler)

	list := v1.Group("/list")
	list.Get("/", DefaulHandler)
	list.Post("/", DefaulHandler)
	list.Patch("/", DefaulHandler)
	list.Delete("/", DefaulHandler)

	membership := v1.Group("/membership")
	membership.Get("/", DefaulHandler)
	membership.Post("/", DefaulHandler)
	membership.Patch("/", DefaulHandler)
	membership.Delete("/", DefaulHandler)

	log.Fatal(app.Listen(":3000"))
}
