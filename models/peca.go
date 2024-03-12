package models

import (
	"errors"
	"regexp"
	"strings"
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

func (p *Peca) Format() error {
	p.Descricao = strings.TrimSpace(p.Descricao)
	if !p.ValidateCod() {
		return errors.New("Codigo invalido")
	}

}

func (p Peca) ValidateCod() bool {
	exp, err := regexp.Compile("([0-9][0-9][.][0-9][0-9][.][0-9][0-9][0-9][0-9])")
	if err != nil {
		return false
	}
	if len(p.Cod) > 10 || len(strings.Split(p.Cod, ".")) > 3 {
		return false
	}
	return exp.Match([]byte(p.Cod))
}
