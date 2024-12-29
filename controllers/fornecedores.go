package controllers

import (
	"api/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type FornecedorRepository interface {
	Create(models.Fornecedor) error
	ListAll() ([]models.Fornecedor, error)
	Find(uint) (models.Fornecedor, error)
	Update(models.Fornecedor) error
	Delete(models.Fornecedor) error
	Search(string) ([]models.Fornecedor, error)
}
type Fornecedor struct {
	r FornecedorRepository
}

func NewFornecedor(r FornecedorRepository) Fornecedor {
	return Fornecedor{r}
}

func (f Fornecedor) CreateFornecedor(c *fiber.Ctx) error {
	var fornecedor models.Fornecedor

	if err := c.BodyParser(&fornecedor); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := fornecedor.Format(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := f.r.Create(fornecedor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fornecedor)
}

func (f Fornecedor) ListAllFornecedores(c *fiber.Ctx) error {
	fornecedores, err := f.r.ListAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fornecedores)
}

func (f Fornecedor) SearchFornecedores(c *fiber.Ctx) error {
	param := c.Query("query")
	if len(param) < 4 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.New("insira ao menos 4 letras na busca").Error()})
	}

	fornecedores, err := f.r.Search(param)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fornecedores)
}

func (f Fornecedor) FindFornecedor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	fornecedor, err := f.r.Find(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(fornecedor)
}

func (f Fornecedor) EditFornecedor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var fornecedor models.Fornecedor

	if err := c.BodyParser(&fornecedor); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}
	fornecedor.ID = uint(id)

	if err := f.r.Update(fornecedor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (f Fornecedor) DeleteForncedor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var fornecedor models.Fornecedor
	fornecedor.ID = uint(id)

	if err := f.r.Delete(fornecedor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
