package controllers

import (
	"api/models"
	"api/repository"

	"github.com/gofiber/fiber/v2"
)

func CreateItemVenda(c *fiber.Ctx) error {
	var item models.ItemVenda
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"erro": err.Error()})
	}

	repo, err := repository.NewItemVendaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON((fiber.Map{"erro": err.Error()}))
	}

	if err := repo.Create(item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"erro": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(item)

}
