package models

import "gorm.io/gorm"

type Venda struct {
	gorm.Model
	Cliente Cliente `gorm:"foreignKey:ID"`
	Peca    Peca    `gorm:"foreignKey:Codigo"`
}

func (v Venda) TableName() string {
	return "venda"

}
