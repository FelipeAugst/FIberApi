package models

import "gorm.io/gorm"

type Venda struct {
	gorm.Model
	Cliente
	Peca
	PecaID     uint
	ClienteID  uint
	Quantidade uint `json:"quantidade" gorm:"column:quantidade"`
}

func (v Venda) TableName() string {
	return "venda"

}
