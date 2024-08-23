package main

import (
	"fmt"

	
	"github.com/HarshithRajesh/idea1/pkg/database"
	"github.com/HarshithRajesh/idea1/pkg/routes"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){
//Database connection
	_,err := database.ConnectionDB()
	if err != nil {
		panic("couldnt connect to db")
	}
	fmt.Println("Connection is successful")
//Intialize fiber app
	app := fiber.New()
//Ading cors middleware 
app.Use(cors.New(cors.Config{
    AllowOrigins:     "https://example.com, https://anotherdomain.com", // Provide a comma-separated string of allowed origins
    AllowMethods:     "GET, POST, PUT, DELETE, PATCH, OPTIONS",         // Provide a comma-separated string of allowed methods
    AllowHeaders:     "Content-Type, Authorization, Accept, Origin",    // Provide a comma-separated string of allowed headers
    AllowCredentials: true,                                             // Only set this to true if you trust the origins listed above
}))

//Routes Setup
	routes.SetUpRoutes(app)

//Start the server

	err = app.Listen(":8000")
	if err != nil {
		panic("Couldnt start the server")
	}
}