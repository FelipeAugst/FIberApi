package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Peca struct {
	gorm.Model
	Descricao string  `json:"descricao" gorm:"column:descricao"`
	Preco     float64 `json:"preco,omitempty" gorm:"column:preco"`
	Custo     float64 `json:"custo,omitempty" gorm:"column:custo"`
	Saldo     uint    `json:"saldo,omitempty" gorm:"column:saldo"`
}

func (p Peca) TableName() string {
	return "peca"
}

func (p *Peca) Format() error {
	p.Descricao = strings.TrimSpace(p.Descricao)
	if err := p.validate(); err != nil {
		return err
	}
	return nil
}

func (p Peca) validate() error {
	if p.Descricao == "" {
		return errors.New("descricao invalida")
	}

	return nil
}
