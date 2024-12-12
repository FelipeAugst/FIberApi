package models

import (
	"gorm.io/gorm"
)

type Venda struct {
	gorm.Model
	ClienteID uint
	Cliente   Cliente     `gorm:"foreignKey:ClienteID"`
	Itens     []ItemVenda `json:"itens"`
}
