package repository

import (
	"api/db"
	"api/models"

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

func (p *peca) List(filter string) ([]models.Peca, error) {

	var pecas []models.Peca
	if err := p.db.Where("descricao like ?", filter+"%").Find(&pecas).Error; err != nil {
		return nil, err
	}

	return pecas, nil
}

func (p *peca) ById(id uint) (models.Peca, error) {

	var peca models.Peca
	if err := p.db.Where("codigo = ?", id).Find(&peca).Error; err != nil {
		return models.Peca{}, err
	}

	return peca, nil
}

func (p *peca) Update(peca models.Peca) error {

	return p.db.Where("id=?", peca.ID).Updates(peca).Error
}

func (p *peca) Delete(peca models.Peca) error {
	return p.db.Delete(&peca).Error

}
