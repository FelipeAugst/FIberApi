package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/paemuri/brdoc"
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

func (c Cliente) TableName() string {
	return "cliente"
}
func (c *Cliente) Format() error {

	c.Nome = strings.TrimSpace(c.Nome)
	c.Endereco = strings.TrimSpace(c.Endereco)
	c.CPF = strings.TrimSpace(c.CPF)
	if err := c.validate(); err != nil {
		return err
	}
	return nil

}

func (c *Cliente) validate() error {
	switch {
	case !brdoc.IsCPF(c.CPF):
		{
			return errors.New("insira um CPF Valido")
		}
	case checkmail.ValidateFormat(c.Email) != nil:
		{
			return errors.New("insira um Email Valido")
		}

	case c.Endereco == "":
		{
			{
				return errors.New("insira um Email Valido")
			}
		}
	case c.Telefone == "":
		{
			{
				return errors.New("insira um Email Valido")
			}
		}
	default:
		return nil
	}
}
