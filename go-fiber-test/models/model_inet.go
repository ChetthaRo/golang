package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type UserData struct {
	Email       string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	NameUser    string `json:"username" validate:"required,min=3,max=32"`
	Pass        string `json:"pass" validate:"required,min=6,max=20"`
	LineId      string `json:"lineid" validate:"min=3,max=32"`
	PhoneNumber string `json:"phonenumber" validate:"required,numeric,len=10"`
	Type        string `json:"type" validate:"required,oneof=retail service"`
	WebSite     string `json:"wepsite" validate:"required,min=2,max=30"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data       []DogsRes `json:"data"`
	Name       string    `json:"name"`
	Count      int       `json:"count"`
	SumRed     int       `json:"countred"`
	SumGree    int       `json:"countgreen"`
	SumpPrink  int       `json:"countprink"`
	SumNoColor int       `json:"countncolor"`
}

type Company struct {
	gorm.Model
	NameCompany string `json:"namecompany"`
	Type        string `json:"type"`
	PhoneNumber string `json:"phonenumber" validate:"required,numeric,len=10"`
	WebSite     string `json:"wepsite" validate:"required,min=2,max=30"`
	ProVince    string `json:"province"`
	SubDisTrict string `json:"subdistrict"`
	DisTrict    string `json:"disTrict"`
	ZipCode     string `json:"zipcode"`
}

type Profile struct {
	gorm.Model
	EmployeeId  int       `json:"employee_id"`
	Name        string    `json:"name"`
	LastName    string    `json:"lastname" `
	BirthdayStr string    `json:"birthdaystr"`
	Birthday    time.Time `json:"-" `
	Age         int       `json:"age"`
	Email       string    `json:"email"`
	Tel         string    `json:"tel"`
}

type ResultProfile struct {
	Data          []ProfileRes `json:"data"`
	Name          string       `json:"name"`
	Count         int          `json:"count"`
	SumGenZ       int          `json:"sumgenz"`
	SumGenY       int          `json:"sumgeny"`
	SumpGenX      int          `json:"sumgenx"`
	SumBabyBoomer int          `json:"sumbabyboomer"`
	SumGeneration int          `json:"sumgeneration"`
}

type ProfileRes struct {
	Name      string `json:"name"`
	ProfileId int    `json: "profile_id"`
	Age       int    `json:"age"`
	Type      string `json:"type"`
}
