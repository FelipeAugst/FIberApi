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
	return &peca{db}, nil
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

func (p peca) Search(param string) ([]models.Peca, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	var pecas []models.Peca
	if err := db.Find(&models.Peca{}, "where codigo like ? or descricao like ?", param, param).Error; err != nil {
		return nil, err
	}

	return pecas, nil
}

func (p peca) Update(pc models.Peca) error {
	return nil

}
