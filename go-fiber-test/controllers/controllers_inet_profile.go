package controllers

import (
	"go-fiber-test/database"
	m "go-fiber-test/models"
	"log"
	"strings"

	"time"

	"github.com/gofiber/fiber/v2"
)

func GetProfiles(c *fiber.Ctx) error {
	db := database.DBConn
	var profiles []m.Profile

	db.Find(&profiles) //delelete = null
	return c.Status(200).JSON(profiles)
}

func CreateProfile(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var ProfileBody m.Profile
	//Check BodyParser must have
	if err := c.BodyParser(&ProfileBody); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	// dob := getDOB(2011, 4, 2)
	ProfileBody.Birthday = ParseBirthDay(ProfileBody.BirthdayStr)
	log.Println(ProfileBody.Birthday)

	db.Create(&ProfileBody)
	return c.Status(201).JSON(ProfileBody)
}

func UpdateProfile(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var profiles m.Profile
	id := c.Params("id")

	//Check BodyParser must have
	if err := c.BodyParser(&profiles); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&profiles)
	return c.Status(201).JSON(profiles)
}

func RemoveProfile(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var profile m.Profile

	id := c.Params("id")

	//check Id in Params

	result := db.Delete(&profile, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}

func GetProfileJson(c *fiber.Ctx) error {
	db := database.DBConn
	var profiles []m.Profile
	sumgenZ := 0
	sumgenY := 0
	sumpgenX := 0
	sumbabyboomer := 0
	sumgeneration := 0

	db.Find(&profiles)

	var profileres []m.ProfileRes

	for _, v := range profiles { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.Age < 24 {
			typeStr = "GenZ"
			sumgenZ++
		} else if v.Age >= 24 && v.Age <= 41 {
			typeStr = "GenY"
			sumgenY++
		} else if v.Age >= 42 && v.Age <= 56 {
			typeStr = "GenX"
			sumpgenX++
		} else if v.Age >= 57 && v.Age <= 75 {
			typeStr = "Baby Boomer"
			sumbabyboomer++
		} else {
			typeStr = "Generation"
			sumgeneration++
		}

		d := m.ProfileRes{
			Name:      v.Name,       //inet1
			ProfileId: v.EmployeeId, //113
			Age:       v.Age,
			Type:      typeStr, //green
		}
		profileres = append(profileres, d)
		// sumAmount += v.Amount
	}

	r := m.ResultProfile{
		Data:          profileres,
		Name:          "golang-test",
		Count:         len(profiles), //หาผลรวม,
		SumGenZ:       sumgenZ,
		SumGenY:       sumgenY,
		SumpGenX:      sumpgenX,
		SumBabyBoomer: sumbabyboomer,
		SumGeneration: sumgeneration,
	}
	return c.Status(200).JSON(r)
}

func GetProfileBySearch(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var profile []m.Profile

	result := db.Where("employee_id = ? OR name = ? OR last_name = ?", search, search, search).Find(&profile)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&profile)
}

func ParseBirthDay(birthdaystr string) time.Time {
	birthday, err := time.Parse("2006-01-02", birthdaystr)
	if err != nil {
		return time.Time{}
	}
	return birthday
}
