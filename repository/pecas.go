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

func (p *peca) Create(peca models.Peca) error {

	return p.db.Create(&peca).Error

}

func (p *peca) ListAll() ([]models.Peca, error) {

	var peca []models.Peca
	if err := p.db.Find(&peca).Error; err != nil {
		return nil, err
	}
	fmt.Println(peca)
	return peca, nil
}

func (p *peca) List(filter string) ([]models.Peca, error) {

	var pecas []models.Peca
	if err := p.db.Where("descricao like ?", filter+"%").Find(&pecas).Error; err != nil {
		return nil, err
	}

	return pecas, nil
}

func (p *peca) Update(peca models.Peca) error {
	fields := make(map[string]any)
	fields["cod"] = peca.Cod
	fields["descricao"] = peca.Descricao
	fields["preco"] = peca.Preco
	return p.db.Where("id=?", peca.ID).Updates(peca).Error
}

func (p *peca) Delete(peca models.Peca) error {
	return p.db.Delete(&peca).Error

}
