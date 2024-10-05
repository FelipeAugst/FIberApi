package models

import (
	"gorm.io/gorm"
)

type Venda struct {
	gorm.Model
	Cliente
	Peca
	Pecas      []PecaVendida
	ClienteID  uint `gorm:"foreignKey:Cliente.ID"`
	Quantidade uint `json:"quantidade" gorm:"column:quantidade"`
}

func (v Venda) TableName() string {
	return "venda"

}
