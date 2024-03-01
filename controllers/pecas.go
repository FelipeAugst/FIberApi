package controllers

import (
	"api/models"
	"api/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreatePeca(c *fiber.Ctx) error {
	var peca models.Peca
	if err := c.BodyParser(&peca); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := r.Create(peca); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.Status(200).JSON(peca)

}

func ListAllPecas(c *fiber.Ctx) error {
	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})

	}
	pecas, err := r.ListAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})

	}
	return c.JSON(pecas)

}

func ListPecas(c *fiber.Ctx) error {
	param := c.Params("filter")
	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	pecas, err := r.List(param)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.Status(200).JSON(pecas)
}

func EditPeca(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})

	}

	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	var peca models.Peca

	err = c.BodyParser(&peca)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	if err := r.Update(id, peca); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})

	}
	return c.SendStatus(200)
}

func DeletePeca(c *fiber.Ctx) error {
	return c.Send([]byte(fmt.Sprintf("Deletar funcionando! peca: %s", c.Params("desc"))))

}
