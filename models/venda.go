package models

import "gorm.io/gorm"

type Venda struct {
	gorm.Model
	Cliente    Cliente `gorm:"foreignKey:CPF"`
	Peca       Peca    `gorm:"foreignKey:Descricao"`
	Quantidade uint    `json:"quantidade" gorm:"column:quantidade" `
}

func (v Venda) TableName() string {
	return "venda"

}
