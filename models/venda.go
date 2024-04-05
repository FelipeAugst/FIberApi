package models

import "gorm.io/gorm"

type Venda struct {
	gorm.Model
	Peca       `gorm:"foreignKey:peca.id"`
	Cliente    `gorm:"foreignKey:cliente.id"`
	PecaID     uint
	ClienteID  uint
	Quantidade uint `json:"quantidade" gorm:"column:quantidade"`
}

func (v Venda) TableName() string {
	return "venda"

}
