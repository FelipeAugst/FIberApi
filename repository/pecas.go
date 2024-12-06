package repository

import (
	"api/db"
	"api/models"

	"gorm.io/gorm"
)

type PecaRepo[T any] interface {
	Create(T) error
	Find(uint) (T, error)
	ListAll() ([]T, error)
	Update(T) error
	Delete(T) error
}

type peca struct {
	db *gorm.DB
}

func NewPecaRepo() (PecaRepo[models.Peca], error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	return &peca{db.Model(&models.Peca{})}, nil
}

func (p *peca) Create(peca models.Peca) error {

	return p.db.Create(&peca).Error

}

func (p *peca) ListAll() ([]models.Peca, error) {

	var peca []models.Peca
	if err := p.db.Find(&peca).Error; err != nil {
		return nil, err
	}

	return peca, nil
}

func (p *peca) Find(id uint) (models.Peca, error) {

	var search models.Peca
	if err := p.db.Find(&search, "id=?", id).Error; err != nil {
		return models.Peca{}, err
	}

	return search, nil
}

func (p *peca) Update(peca models.Peca) error {

	return p.db.Updates(peca).Error
}

func (p *peca) Delete(peca models.Peca) error {
	return p.db.Delete(&peca).Error

}
