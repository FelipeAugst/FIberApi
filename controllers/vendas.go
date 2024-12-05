package controllers

import (
	"api/models"
	"api/repository"

	"github.com/gofiber/fiber/v2"
)

func CreateVenda(c *fiber.Ctx) error {
	var venda models.Venda
	if err := c.BodyParser(&venda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	vr, err := repository.NewVendaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	if err := vr.Create(venda); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())

	}

	return c.Status(fiber.StatusCreated).JSON(venda)

}
