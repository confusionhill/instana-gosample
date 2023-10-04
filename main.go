package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	instana "github.com/instana/go-sensor"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	instana.NewSensor("example-go")
	app := fiber.New()

	// Define a route handler for the root path "/"
	app.Get("/", func(c *fiber.Ctx) error {
		log.Error(errors.New("hello world!"))
		return c.SendString("Hello, Fiber!")
	})

	// Start the Fiber application on port 3000
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
