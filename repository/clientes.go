package repository

import (
	"api/db"
	"api/models"

	"gorm.io/gorm"
)

type cliente struct {
	db *gorm.DB
}

func NewClienteRepo() (repository[models.Cliente], error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	return &cliente{db.Model(&models.Cliente{})}, nil

}

func (c *cliente) Create(cliente models.Cliente) error {

	return c.db.Create(&cliente).Error

}

func (c *cliente) ListAll() ([]models.Cliente, error) {
	var cliente []models.Cliente
	err := c.db.Find(&cliente).Error
	if err != nil {
		return nil, err
	}
	return cliente, nil

}

func (c *cliente) List(param string) ([]models.Cliente, error) {
	var cliente []models.Cliente
	if err := c.db.Where("nome like ?", "%"+param).Find(&cliente).Error; err != nil {
		return nil, err
	}
	return cliente, nil
}

func (c *cliente) Update(cliente models.Cliente) error {

	return c.db.Where("id= ?", cliente.ID).Updates(cliente).Error

}

func (c *cliente) Delete(cliente models.Cliente) error {
	return c.db.Delete(&cliente).Error

}
