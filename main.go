package main

import (
	"Crawl_Web/configs"

	"Crawl_Web/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//download from Web
	configs.DownloadURL()

	//routes
	routes.DomainRoute(app)

	app.Listen(":8080")
}
