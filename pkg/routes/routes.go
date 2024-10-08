package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/HarshithRajesh/idea1/pkg/controllers"
)

func SetUpRoutes(app *fiber.App){

	app.Get("/",controllers.Hello)
	app.Post("/api/register",controllers.Register)
	app.Post("/api/login",controllers.Login)
	app.Get("/api/user",controllers.User)
	app.Post("/logout",controllers.Logout)
}