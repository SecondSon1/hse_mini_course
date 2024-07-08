package main

import (
	"hse_mini_course/accounts"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	PORT = ":6969"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()
	app.Use(logger.New())

	accountHandler := accounts.NewHandler()

	app.Post("/account/new", accountHandler.CreateAccount)

	group := app.Group("/account/:name")
	group.Get("/", accountHandler.GetUser)
	group.Post("/", accountHandler.NewTransaction)
	group.Patch("/", accountHandler.ChangeName)
	group.Delete("/", accountHandler.DeleteAccount)

	log.Fatal(app.Listen(PORT))
}
