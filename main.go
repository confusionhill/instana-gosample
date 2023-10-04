package main

import (
	"github.com/gofiber/fiber/v2"
	instana "github.com/instana/go-sensor"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	app := fiber.New()

	opt := *instana.DefaultOptions()
	opt.Service = "example-golang-fiber-instana"
	opt.EnableAutoProfile = true
	instana.StartMetrics(&opt)
	// initialize and configure the logger
	logger := logrus.New()
	logger.Level = logrus.InfoLevel

	// check if INSTANA_DEBUG is set and set the log level to DEBUG if needed
	if _, ok := os.LookupEnv("INSTANA_DEBUG"); ok {
		logger.Level = logrus.DebugLevel
	}

	// use logrus to log the Instana Go Collector messages
	instana.SetLogger(logger)

	// Define a route handler for the root path "/"
	app.Get("/", func(c *fiber.Ctx) error {
		logger.Error("hahahaha hacker!")
		return c.SendString("Hello, Fiber!")
	})

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		logger.Error("hacker mana lagi nih?")
		return ctx.SendString("Emg boleh begitu?")
	})

	// Start the Fiber application on port 3000
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
