package models

import (
	"gorm.io/gorm"
)

type Venda struct {
	gorm.Model
	Peca            Peca    `json:"pecas"`
	Cliente         Cliente `json:"cliente"`
	PecaID          uint
	ClienteID       uint
	PedidoDeVendaID uint
}

func (v Venda) TableName() string {
	return "venda"

}

type PedidoDeVenda struct {
	gorm.Model
	Vendas []Venda `json:"vendas"`
}
