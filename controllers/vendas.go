package controllers

import (
	"api/models"
	"api/repository"

	"github.com/gofiber/fiber/v2"
)

func CreateVenda(c *fiber.Ctx) error {
	idCliente, err := c.ParamsInt("cliente")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	idPeca, err := c.ParamsInt("peca")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	rc, err := repository.NewClienteRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	cliente, err := rc.ById(uint(idCliente))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	rp, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	peca, err := rp.ById(uint(idPeca))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	q := make(map[string]uint)

	if err := c.BodyParser(&q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if q["quantidade"] <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Quantidade Invalida"})
	}

	if peca.Saldo < q["quantidade"] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Saldo insuficiente"})
	}

	var venda = models.Venda{
		Cliente:    cliente,
		Peca:       peca,
		Quantidade: q["quantidade"],
	}

	rv, err := repository.NewVendaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := rv.Create(venda); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(venda)

}
func ListVendas(ctx *fiber.Ctx) error {

	r, err := repository.NewVendaRepo()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	vendas, err := r.ListAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(vendas)

}

func DeleteVendas(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var venda models.Venda
	venda.ID = uint(id)

	r, err := repository.NewVendaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := r.Delete(venda.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)

}
