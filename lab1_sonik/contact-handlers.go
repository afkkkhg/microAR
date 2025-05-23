package main

import "github.com/gofiber/fiber/v2"

// Класс телефона, который вложен в класс Контакт с доп данными
type Phone struct {
	TypeID      int `json:"type_id,omitempty"`
	CountryCode int `json:"country_code,omitempty"`
	Operator    int `json:"operator,omitempty"`
	Number      int `json:"number,omitempty"`
}

// Класс контакт с полями, содержащий
type Contact struct {
	ID         int      `json:"id,omitempty"`
	Username   string   `json:"username,omitempty"`
	GivenName  string   `json:"given_name,omitempty"`
	FamilyName string   `json:"family_name,omitempty"`
	Phones     []Phone  `json:"phones,omitempty"`
	Emails     []string `json:"emails,omitempty"`
	Birthdate  string   `json:"birthdate,omitempty"`
}

// создание экземпляра из класса с пустой структурой, чтобы не прописывать все данные (они там есть, но пустые)
func NewContactNoArg() *Contact {
	return &Contact{}
}

func ContactCreate(ctx *fiber.Ctx) error {
	var contact Contact
// условие для считывания тела запроса и заполнение полей для переменной
	if err := ctx.BodyParser(&contact); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(contact)
}

func ContactRead(ctx *fiber.Ctx) error {
	idUser := ctx.QueryInt("id")
	testContactNoArg := NewContactNoArg()
	testContactNoArg.ID = idUser
	return ctx.JSON(testContactNoArg)
}

func ContactUpdate(ctx *fiber.Ctx) error {
	idUser := ctx.QueryInt("id")
	testContactNoArg := NewContactNoArg()
	testContactNoArg.ID = idUser
	return ctx.JSON(testContactNoArg)
}

func ContactDelete(ctx *fiber.Ctx) error {
	idUser := ctx.QueryInt("id")
	testContactNoArg := NewContactNoArg()
	testContactNoArg.ID = idUser
	return ctx.JSON(testContactNoArg)
}
