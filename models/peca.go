package models

import (
	"errors"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type Peca struct {
	gorm.Model
	Cod       string  `json:"cod" gorm:"column:cod;type=int;unique"`
	Descricao string  `json:"descricao" gorm:"column:descricao"`
	Preco     float64 `json:"preco" gorm:"column:preco"`
	Custo     float64 `json:"custo" gorm:"column:custo"`
	Saldo     float64 `json:"saldo" gorm:"column:saldo"`
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
	exp, err := regexp.Compile("([0-9][0-9][.][0-9][0-9][.][0-9][0-9][0-9][0-9])")
	if err != nil {
		return err
	}
	if len(p.Cod) > 10 || len(strings.Split(p.Cod, ".")) > 3 || !exp.Match([]byte(p.Cod)) {
		return errors.New("padrao invalido para o codigo da peca")
	}
	return nil
}
