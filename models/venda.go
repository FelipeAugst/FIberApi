package models

import (
	"gorm.io/gorm"
)

type Venda struct {
	gorm.Model
	Cliente
	Peca
	Pecas     []PecaVendida `json:"pecas"`
	ClienteID uint          `gorm:"foreignKey:Cliente.ID"`
}

func (v Venda) TableName() string {
	return "venda"

}
