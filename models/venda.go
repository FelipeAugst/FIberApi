package models

import (
	"gorm.io/gorm"
)

type ItemVenda struct {
	gorm.Model
	PecaID     uint
	Peca       Peca `json:"pecas" gorm:"foreignKey:PecaID"`
	Quantidade int  `json:"quantidade"`
	VendaID    uint
	Venda      Venda `gorm:"foreignKey:VendaID"`
}

func (i ItemVenda) TableName() string {
	return "Item"

}

type Venda struct {
	gorm.Model
	ClienteID uint
	Cliente   Cliente     `gorm:"foreignKey:ClienteID"`
	Itens     []ItemVenda `json:"itens"`
}
