package repository

import (
	"api/db"
	"api/models"

	"gorm.io/gorm"
)

type ClienteRepository[T any] interface {
	Create(T) error
	GetAll() ([]T, error)
	Get(id uint) (T, error)
	Update(T) error
	Delete(T) error
}

type cliente struct {
	db *gorm.DB
}

func NewClienteRepo() (ClienteRepository[models.Cliente], error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	return &cliente{db.Model(&models.Cliente{})}, nil

}

func (c *cliente) Create(cliente models.Cliente) error {

	return c.db.Create(&cliente).Error

}

func (c *cliente) GetAll() ([]models.Cliente, error) {
	var cliente []models.Cliente
	err := c.db.Find(&cliente).Error
	if err != nil {
		return nil, err
	}
	return cliente, nil

}

func (c *cliente) Get(id uint) (models.Cliente, error) {
	var SearchedClient models.Cliente
	SearchedClient.ID = id

	if err := c.db.Find(&SearchedClient).Error; err != nil {
		return models.Cliente{}, err
	}
	return SearchedClient, nil
}

func (c *cliente) Update(cliente models.Cliente) error {

	return c.db.Where("id= ?", cliente.ID).Updates(cliente).Error

}

func (c *cliente) Delete(cliente models.Cliente) error {
	return c.db.Delete(&cliente).Error

}
