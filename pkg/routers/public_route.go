package routers

import (
	"Template/pkg/controllers"
	"Template/pkg/controllers/healthchecks"

	"github.com/gofiber/fiber/v2"
)

func SetupPublicRoutes(app *fiber.App) {

	// Endpoints
	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/public")
	v1Endpoint := publicEndpoint.Group("/v1")

	testRoutes := v1Endpoint.Group("/test")
	testRoutes.Get("/convert", controllers.XmltoJson)
	testRoutes.Get("/read", controllers.Consumer_Kafka)
	testRoutes.Get("/sql", controllers.MySQL_Read)

	// Service health check
	v1Endpoint.Get("/", healthchecks.CheckServiceHealth)
}

func SetupPublicRoutesB(app *fiber.App) {

	// Endpoints
	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/public")
	v1Endpoint := publicEndpoint.Group("/v1")

	// Service health check
	v1Endpoint.Get("/", healthchecks.CheckServiceHealthB)
}
