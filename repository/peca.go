package repository

import (
	"api/db"
	"api/models"
	"fmt"

	"gorm.io/gorm"
)

type peca struct {
	db *gorm.DB
}

func (p peca) TableName() string {
	return "peca"
}
func NewPecaRepo() (repository[models.Peca], error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	return &peca{db.Model(&models.Peca{})}, nil
}

func (p peca) Create(peca models.Peca) error {

	return p.db.Create(&peca).Error

}

func (p peca) ListAll() ([]models.Peca, error) {

	var peca []models.Peca
	if err := p.db.Find(&peca).Error; err != nil {
		return nil, err
	}
	fmt.Println(peca)
	return peca, nil
}

func (p peca) List(filter string) ([]models.Peca, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	var pecas []models.Peca
	if err := db.Find(&pecas, "descricao like ?%", filter).Error; err != nil {
		return nil, err
	}

	return pecas, nil
}

func (p peca) Update(ID int, peca models.Peca) error {
	return p.db.Save(&peca).Where(ID).Error
}

func (p peca) Delete(ID int) error {
	return p.db.Delete(ID).Error

}
