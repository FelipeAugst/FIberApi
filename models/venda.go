package models

import "gorm.io/gorm"

type Venda struct {
	gorm.Model
	Cliente Cliente `gorm:"foreignKey:CPF"`
	Peca    Peca    `gorm:"foreignKey:Cod"`
}

func (v Venda) TableName() string {
	return "venda"

}
