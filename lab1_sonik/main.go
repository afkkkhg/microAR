package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// адреса для групп
	app.Post("/api/v1/group", GroupCreate)
	app.Get("/api/v1/group", GroupRead)
	app.Put("/api/v1/group", GroupChange)
	app.Delete("/api/v1/group", GroupDelete)
	// адрес для контактов
	app.Get("api/v1/contact", ContactRead)
	app.Post("api/v1/contact", ContactCreate)
	app.Put("api/v1/contact", ContactUpdate)
	app.Delete("api/v1/contact", ContactDelete)
	// порт сервера
	if err := app.Listen(":6080"); err != nil {
		log.Fatal(err)
	}
}
