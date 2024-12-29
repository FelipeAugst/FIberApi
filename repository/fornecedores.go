package repository

import (
	"gorm.io/gorm"

	"api/db"
	"api/models"
)

type FornecedorRepo[T any] interface {
	Create(T) error
	ListAll() ([]T, error)
	Find(id uint) (T, error)
	Update(T) error
	Delete(T) error
	Search(string) ([]models.Fornecedor, error)
}

type Fornecedor struct {
	db *gorm.DB
}

func NewFornecedorRepo() (Fornecedor, error) {
	db, err := db.Connect()
	if err != nil {
		return Fornecedor{}, err
	}
	return Fornecedor{db.Model(&models.Fornecedor{})}, nil
}

func (f *Fornecedor) Create(fornecedor models.Fornecedor) error {
	return f.db.Create(&fornecedor).Error
}

func (f *Fornecedor) Find(id uint) (models.Fornecedor, error) {
	var fornecedor models.Fornecedor
	if err := f.db.Where("ID = ?", id).Find(&fornecedor).Error; err != nil {
		return models.Fornecedor{}, err
	}

	return fornecedor, nil
}

func (f *Fornecedor) ListAll() ([]models.Fornecedor, error) {
	var fornecedores []models.Fornecedor
	if err := f.db.Find(&fornecedores).Error; err != nil {
		return nil, err
	}

	return fornecedores, nil
}

func (f *Fornecedor) Search(filter string) ([]models.Fornecedor, error) {
	var fornecedores []models.Fornecedor
	if err := f.db.Find(&fornecedores, "NOME LIKE ?", filter+"%").Error; err != nil {
		return nil, err
	}

	return fornecedores, nil
}

func (f *Fornecedor) Update(fornecedor models.Fornecedor) error {
	return f.db.Where("id=?", fornecedor.ID).Updates(fornecedor).Error
}

func (f *Fornecedor) Delete(fornecedor models.Fornecedor) error {
	return f.db.Delete(&fornecedor).Error
}
