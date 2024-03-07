package models

import (
	"errors"
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

func (c *Cliente) Format() error {

	c.Nome = strings.TrimSpace(c.Nome)
	c.Endereco = strings.TrimSpace(c.Endereco)
	c.CPF = strings.TrimSpace(c.CPF)
	return nil

}

func (c *Cliente) validate() error {
	switch {
	case c.CPF == "":
		{
			return errors.New("insira um CPF Valido")
		}
	case c.Email == "":
		{
			return errors.New("insira um Email Valido")
		}

	}
	return nil
}
