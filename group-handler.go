package main

import (
	"github.com/gofiber/fiber/v2"
)

type Group struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Contacts    []int  `json:"contacts,omitempty"`
}

func NewGroupNoArg() *Group {
	return &Group{}
}

func GroupCreate(ctx *fiber.Ctx) error {
	var group Group
	if err := ctx.BodyParser(&group); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.JSON(group)
}

func GroupRead(ctx *fiber.Ctx) error {
	idGroup := ctx.QueryInt("id")
	testGroupNoArg := NewGroupNoArg()
	testGroupNoArg.ID = idGroup
	return ctx.JSON(testGroupNoArg)
}

func GroupChange(ctx *fiber.Ctx) error {
	idGroup := ctx.QueryInt("id")
	testGroupNoArg := NewGroupNoArg()
	testGroupNoArg.ID = idGroup
	return ctx.JSON(testGroupNoArg)
}
func GroupDelete(ctx *fiber.Ctx) error {
	idGroup := ctx.QueryInt("id")
	testGroupNoArg := NewGroupNoArg()
	testGroupNoArg.ID = idGroup
	return ctx.JSON(testGroupNoArg)
}
