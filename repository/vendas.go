package repository

import (
	"api/db"
	"api/models"

	"gorm.io/gorm"
)

type VendaRepo[T any] interface {
	Create(T) error
	Find(uint) (T, error)
	ListAll() ([]T, error)
	Update(T) error
	Delete(T) error
	Conclude(id uint) error
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

func (v *venda) Create(m models.Venda) error {
	m.Concluida = false
	return v.db.Create(&m).Error

}

func (v *venda) Find(id uint) (models.Venda, error) {
	var search models.Venda

	if err := v.db.Preload("Clientes").Preload("Itemvenda").Find(search, "id=?", id).Error; err != nil {
		return models.Venda{}, err
	}
	return search, nil
}

func (v *venda) ListAll() ([]models.Venda, error) {
	var vendas []models.Venda
	if err := v.db.Preload("Clientes", "clientes.id=vendas.clienteid").Find(&vendas).Error; err != nil {
		return nil, err
	}
	return vendas, nil

}

func (v *venda) Update(updated models.Venda) error {
	return v.db.Where("id=?", updated.ID).Updates(updated).Error

}

func (v *venda) Delete(deleted models.Venda) error {
	return v.db.Delete(&deleted).Error

}

func (v *venda) Conclude(id uint) error {

	return v.db.Where("ID = ?", id).UpdateColumn("Concluida", true).Error

}
