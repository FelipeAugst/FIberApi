package models

import "gorm.io/gorm"

type ItemVenda struct {
	gorm.Model
	PecaID     uint
	Peca       Peca `json:"pecas" gorm:"foreignKey:PecaID"`
	Quantidade int  `json:"quantidade"`
	VendaID    uint
	Venda      Venda `gorm:"foreignKey:VendaID"`
}
