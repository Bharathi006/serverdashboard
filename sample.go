package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func main() {
	// creating fiber app

	//load template
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	//Configure middleware

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Create a logger file
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer file.Close()
	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	// static file path
	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Name": "Welcome to kiriyatech",
		})

	})

	app.Post("/", func(c *fiber.Ctx) error {
		var body struct {
			Message string
		}
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		return c.Render("index", fiber.Map{
			"Name":    "Welcome kiriyatech",
			"Message": body.Message,
		})
	})

	app.Listen(":8080")

}
