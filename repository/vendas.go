package repository

import (
	"api/db"
	"api/models"

	"gorm.io/gorm"
)

type VendaRepo[T any] interface {
	Create(T) error
	ListAll() ([]T, error)
	Update(uint) error
	Delete(uint) error
	ByAgent(uint) ([]T, error)
	ByProduct(uint) ([]T, error)
}

type venda struct {
	db *gorm.DB
}

func NewVendaRepo() (VendaRepo[models.Venda], error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	return &venda{db.Model(&models.Venda{})}, nil
}

func (v venda) Create(venda models.Venda) error {
	venda.Peca.Saldo -= venda.Quantidade
	if err := v.db.Model(&models.Peca{}).Raw("update peca set saldo=? where id=?", venda.Peca.Saldo, venda.Peca.ID).Error; err != nil {
		return err
	}
	return v.db.Preload("Cliente").Preload("Peca").Create(&venda).Error

}

func (v venda) ListAll() ([]models.Venda, error) {
	var vendas []models.Venda
	if err := v.db.Model(&models.Venda{}).Preload("Cliente").Preload("Peca").Find(&vendas).Error; err != nil {
		return nil, err
	}

	return vendas, nil

}

func (v venda) ByAgent(ID uint) ([]models.Venda, error) {
	var cliente models.Cliente
	cliente.ID = ID
	var vendas []models.Venda
	if err := v.db.Preload("Cliente").Preload("Peca").Where("Clliente=?", cliente.ID).Find(&vendas).Error; err != nil {
		return nil, err
	}

	return vendas, nil

}

func (v venda) ByProduct(ID uint) ([]models.Venda, error) {
	var peca models.Peca
	peca.ID = ID
	var vendas []models.Venda
	if err := v.db.Preload("Cliente").Preload("Peca").Where("Peca=?", peca.ID).Find(&vendas).Error; err != nil {
		return nil, err
	}

	return vendas, nil

}

func (v venda) Delete(ID uint) error {
	var venda models.Venda
	venda.ID = ID

	err := v.db.Find(&venda).Error
	if err != nil {
		return err
	}

	if err := v.db.Model(&models.Peca{}).Raw("update peca set saldo=saldo+? where id=?", venda.Quantidade, venda.Peca.ID).Error; err != nil {
		return err
	}
	return v.db.Where("id = ?", venda.ID).Delete(venda).Error

}

func (v venda) Update(ID uint) error {
	var venda models.Venda
	venda.ID = ID
	return v.db.Where("id = ?", venda.ID).Updates(venda).Error

}
