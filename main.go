package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	instana "github.com/instana/go-sensor"
	"github.com/instana/go-sensor/instrumentation/instafiber"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	opt := *instana.DefaultOptions()
	opt.Service = "example-golang-fiber-instana"
	opt.EnableAutoProfile = true
	sensor := instana.NewSensorWithTracer(instana.NewTracerWithOptions(&opt))
	// initialize and configure the logger
	logger := logrus.New()
	logger.Level = logrus.InfoLevel

	// check if INSTANA_DEBUG is set and set the log level to DEBUG if needed
	if _, ok := os.LookupEnv("INSTANA_DEBUG"); ok {
		logger.Level = logrus.DebugLevel
	}

	// use logrus to log the Instana Go Collector messages
	instana.SetLogger(logger)

	app := fiber.New()

	app.Get("/greet", instafiber.TraceHandler(sensor, "greet", "/greet", func(c *fiber.Ctx) error {
		_, err := http.Get("http://119.81.37.230:1323/place/something")
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString("Hello world!")
	}))

	app.Get("/login", instafiber.TraceHandler(sensor, "login", "/login", func(ctx *fiber.Ctx) error {
		jsonStr := []byte(`{"email": "jamet@kudasi.com", "password": "kolorijo123"}`)
		// Create a new HTTP request with the JSON body
		req, err := http.NewRequest("POST", "http://119.81.37.230:1323/login", bytes.NewBuffer(jsonStr))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return err
		}

		// Set the request headers
		req.Header.Set("Content-Type", "application/json")

		// Create an HTTP client and send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return err
		}
		defer resp.Body.Close()

		result := map[string]interface{}{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return err
		}
		return ctx.JSON(result)
	}))

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
	err := app.Listen(":3001")
	if err != nil {
		panic(err)
	}
}
