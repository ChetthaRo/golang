package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")

	v1.Get("/", c.HelloTestV1)
	v1.Get("/user/:name", c.ParamsTest)

	v2.Get("/", c.HelloTestV2)

	v3 := api.Group("/v3")

	pond := v3.Group("/pond")

	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/de", c.GetDogsDelete)
	dog.Get("/json", c.GetDogsJson)
	dog.Get("/hp", c.GetDogsHp)

	company := v1.Group("/company")
	company.Get("", c.GetCompanys)

	//CRUD Profile
	profile := v1.Group("/profile")
	profile.Get("", c.GetProfiles)
	profile.Get("/user/json", c.GetProfileJson)
	profile.Get("/user/filter", c.GetProfileBySearch)

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	}))

	// จัดกรุ๊ป

	//get data from params
	//sent data from body
	v1.Post("/", c.BodyParserTest)
	//post data from queryParams
	v1.Post("/inet", c.QueryTest)
	v1.Post("/valid", c.ValidateTest)
	v1.Post("/fact/:number", c.FactTest)

	pond.Post("/", c.AsciCode)

	v1.Post("/register", c.Register)

	dog.Post("", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)

	// company.Get("/filter", c.Getcompany)
	// company.Get("/json", c.GetcompanysJson)
	company.Post("", c.CreateCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)
	profile.Post("/user", c.CreateProfile)
	profile.Put("/user/:id", c.UpdateProfile)
	profile.Delete("/user/:id", c.RemoveProfile)

}
