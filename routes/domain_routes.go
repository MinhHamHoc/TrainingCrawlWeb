package routes

import (
	"github.com/gofiber/fiber/v2"

	"Crawl_Web/controllers"
)

func DomainRoute(app *fiber.App) {
	app.Get("/domain", controllers.ReadAllDomain)
	app.Get("/domain/:domainId", controllers.GetADomain)
	app.Post("/domain/create", controllers.CreateADomain)
	app.Delete("domain/:domainId", controllers.DeleteADomain)
	app.Put("domain/:domainId", controllers.UpdateADomain)

}
