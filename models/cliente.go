package models

import (
	"strings"

	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	Nome     string `json:"nome" gorm:"column:nome;type:text"`
	Endereco string `json:"endereco" gorm:"column:endereco;type:text"`
	Email    string `json:"email" gorm:"column:email;type:varchar(50)"`
	CPF      string `json:"cpf" gorm:"column:cpf;type:text"`
	Telefone string `json:"telefone" gorm:"column:telefone;type:text"`
}

func (c *Cliente) Format() {

	c.Nome = strings.TrimSpace(c.Nome)
	c.Endereco = strings.TrimSpace(c.Endereco)
	c.CPF = strings.TrimSpace(c.CPF)

}
