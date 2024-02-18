package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	// จัดกรุ๊ป
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", c.HelloTestV1)
	//get data from params
	v1.Get("/user/:name", c.ParamsTest)
	//sent data from body
	v1.Post("/", c.BodyParserTest)
	//post data from queryParams
	v1.Post("/inet", c.QueryTest)
	v1.Post("/valid", c.ValidateTest)
	v1.Post("/fact/:number", c.FactTest)

	v2 := api.Group("/v2")
	v2.Get("/", c.HelloTestV2)

	v3 := api.Group("/v3")
	pond := v3.Group("/pond")
	pond.Post("/", c.AsciCode)

	v1.Post("/register", c.Register)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/de", c.GetDogsDelete)
	dog.Get("/json", c.GetDogsJson)
	dog.Get("/hp", c.GetDogsHp)
	dog.Post("", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)

	//CRUD Company
	company := v1.Group("/company")
	company.Get("", c.GetCompanys)
	// company.Get("/filter", c.Getcompany)
	// company.Get("/json", c.GetcompanysJson)
	company.Post("", c.CreateCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)

}
