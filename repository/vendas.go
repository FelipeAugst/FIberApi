package repository

import (
	"api/db"
	"api/models"

	"gorm.io/gorm"
)

type venda struct {
	db *gorm.DB
}

func NewVendaRepo() (operation[models.Venda], error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	return &venda{db.Model(&models.Venda{})}, nil
}

func (v venda) Create(venda models.Venda) error {
	return v.db.Create(venda).Error

}

func (v venda) ListAll() ([]models.Venda, error) {
	var vendas []models.Venda
	if err := v.db.Find(&vendas).Error; err != nil {
		return nil, err
	}

	return vendas, nil

}

func (v venda) ByAgent(ID uint) ([]models.Venda, error) {
	var cliente models.Cliente
	cliente.ID = ID
	var vendas []models.Venda
	if err := v.db.Where(cliente).Find(&vendas).Error; err != nil {
		return nil, err
	}

	return vendas, nil

}

func (v venda) ByProduct(ID uint) ([]models.Venda, error) {
	var peca models.Peca
	peca.ID = ID
	var vendas []models.Venda
	if err := v.db.Where(peca).Find(&vendas).Error; err != nil {
		return nil, err
	}

	return vendas, nil

}

func (v venda) Delete(ID uint) error {
	var venda models.Venda
	venda.ID = ID
	return v.db.Where("id = ?", venda.ID).Delete(venda).Error

}

func (v venda) Update(ID uint) error {
	var venda models.Venda
	venda.ID = ID
	return v.db.Where("id = ?", venda.ID).Updates(venda).Error

}
