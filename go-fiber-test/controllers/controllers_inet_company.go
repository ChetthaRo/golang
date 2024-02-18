package controllers

import (
	"go-fiber-test/database"
	m "go-fiber-test/models"

	"github.com/gofiber/fiber/v2"
)

func GetCompanys(c *fiber.Ctx) error {
	// db := database.DBConn
	// var dogs []m.Dogs
	db := database.DBConn
	var companys []m.Company

	db.Find(&companys) //delelete = null
	return c.Status(200).JSON(companys)
}

func CreateCompany(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var companyBody m.Company

	//Check BodyParser must have
	if err := c.BodyParser(&companyBody); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&companyBody)
	return c.Status(201).JSON(companyBody)
}

func UpdateCompany(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	//Check BodyParser must have
	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)
	return c.Status(201).JSON(company)
}

func RemoveCompany(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var company m.Company

	id := c.Params("id")

	//check Id in Params

	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
