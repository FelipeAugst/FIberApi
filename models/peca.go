package models

import (
	"time"

	"gorm.io/gorm"
)

type Peca struct {
	gorm.Model
	Cod       string    `json:"cod" gorm:"column:cod;type=int"`
	Descricao string    `json:"descricao" gorm:"column:descricao"`
	Preco     float64   `json:"preco" gorm:"column:preco"`
	LeadTime  time.Time `json:"leadtime" gorm:"column:leadtime"`
}

func (p Peca) TableName() string {
	return "peca"
}
