package controllers

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"go-fiber-test/database"
	m "go-fiber-test/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HelloTestV1(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func HelloTestV2(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func BodyParserTest(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error {

	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is  " + a
	return c.JSON(str)
}

func ValidateTest(c *fiber.Ctx) error {

	//Connect to database

	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

func FactTest(c *fiber.Ctx) error {
	numberStr := c.Params("number")
	number, err := strconv.Atoi(numberStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Number")
	}

	reusult := Factorial(number)
	res := fmt.Sprintf(" %d! = %d ", number, reusult)
	return c.SendString(res)
}

func AsciCode(c *fiber.Ctx) error {
	taxId := c.Query("tax_id")
	if taxId == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Plase Enter tax_id",
		})
	}
	var result []int
	for i := 0; i < len(taxId); i++ {
		result = append(result, int(taxId[i]))
	}

	return c.JSON(result)
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Register(c *fiber.Ctx) error {
	user := new(m.UserData)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	// userData := user.UserName

	regex := regexp.MustCompile(`^[A-Za-z0-9_-]+$`)
	dataUsername := regex.MatchString(user.NameUser) && strings.ContainsAny(user.NameUser, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") && strings.ContainsAny(user.NameUser, "abcdefghijklmnopqrstuvwxyz") && strings.ContainsAny(user.NameUser, "0123456789") && strings.ContainsAny(user.NameUser, "-_")

	if !dataUsername {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Username Format",
		})
	}

	return c.JSON(user)
}

func DogIDGreaterThan100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ?", 100)
}
func DogID(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ? AND dog_id < 100", 50)
}

func GetDogsHp(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Scopes(DogID).Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs
	countRed := 0
	countGreen := 0
	countPrink := 0
	countNocolor := 0

	db.Find(&dogs)

	var dogsr []m.DogsRes

	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID > 50 && v.DogID < 100 {
			typeStr = "red"
			countRed++
		} else if v.DogID > 113 && v.DogID < 150 {
			typeStr = "green"
			countGreen++
		} else if v.DogID > 200 && v.DogID < 250 {
			typeStr = "pink"
			countPrink++
		} else {
			typeStr = "no color"
			countNocolor++
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet1
			DogID: v.DogID, //113
			Type:  typeStr, //green
		}
		dogsr = append(dogsr, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		Data:       dogsr,
		Name:       "golang-test",
		Count:      len(dogs), //หาผลรวม,
		SumRed:     countRed,
		SumGree:    countGreen,
		SumpPrink:  countPrink,
		SumNoColor: countNocolor,
	}
	return c.Status(200).JSON(r)
}

func GetDogsDelete(c *fiber.Ctx) error {
	db := database.DBConn
	var dog []m.Dogs
	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dog)
	return c.Status(200).JSON(dog)
}
