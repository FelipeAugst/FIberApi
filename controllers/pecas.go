package controllers

import (
	"api/models"
	"api/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CriaPeca(c *fiber.Ctx) error {
	var peca models.Peca
	if err := c.BodyParser(&peca); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}

	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
	}
	if err := r.Create(peca); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
	}

	return c.Status(200).JSON(peca)

}

func ListaPecas(c *fiber.Ctx) error {
	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})

	}
	pecas, err := r.ListAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})

	}
	return c.JSON(pecas)

}

func BuscaPecas(c *fiber.Ctx) error {
	return c.Send([]byte(fmt.Sprintf("Buscar funcionando! peca: %s", c.Params("desc"))))

}

func EditaPeca(c *fiber.Ctx) error {
	return c.Send([]byte(fmt.Sprintf("Editar funcionando! peca: %s", c.Params("desc"))))

}

func DeletaPeca(c *fiber.Ctx) error {
	return c.Send([]byte(fmt.Sprintf("Deletar funcionando! peca: %s", c.Params("desc"))))

}
