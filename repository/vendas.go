package repository

import (
	"api/db"
	"api/models"

	"gorm.io/gorm"
)

type venda struct {
	db *gorm.DB
}

func NewVendaRepo() (*venda, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	return &venda{db.Model(&models.Venda{})}, nil
}

func (v venda) Create(venda models.Venda) error {
	return v.db.Create(venda).Error

}
