package controllers

import (
	"api/models"
	"api/repository"

	"github.com/gofiber/fiber/v2"
)

func CreateCliente(c *fiber.Ctx) error {
	var cliente models.Cliente
	err := c.BodyParser(&cliente)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	r, err := repository.NewClienteRepo()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := r.Create(cliente); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cliente)
}

func ListAllClientes(c *fiber.Ctx) error {
	var clientes []models.Cliente
	r, err := repository.NewClienteRepo()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	clientes, err = r.ListAll()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(clientes)

}

func ListClientes(c *fiber.Ctx) error {

	filter := c.Params("filter")
	var clientes []models.Cliente
	r, err := repository.NewClienteRepo()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	clientes, err = r.List(filter)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(clientes)

}
