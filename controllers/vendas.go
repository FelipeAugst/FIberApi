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

func ListAllVendas(c *fiber.Ctx) error {

	repo, err := repository.NewVendaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	vendas, err := repo.ListAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"erro": err.Error()})
	}
	return c.JSON(vendas)

}

func FindVenda(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"erro": err.Error()})
	}
	repo, err := repository.NewVendaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	venda, err := repo.Find(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(venda)

}

func UpdateVenda(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"erro": err.Error()})
	}
	var venda models.Venda

	if err := c.BodyParser(&venda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"erro": err.Error()})
	}
	venda.ID = uint(id)

	repo, err := repository.NewVendaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := repo.Update(venda); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(venda)

}

func DeleteVenda(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"erro": err.Error()})
	}
	var venda models.Venda

	venda.ID = uint(id)

	repo, err := repository.NewVendaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := repo.Delete(venda); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)

}

func ConcludeVenda(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"erro": err.Error()})
	}

	repo, err := repository.NewVendaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"erro": err.Error()})
	}

	if err := repo.Conclude(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"erro": err.Error()})

	}
	return c.SendStatus(fiber.StatusOK)
}
