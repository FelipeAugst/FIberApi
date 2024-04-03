package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/paemuri/brdoc"
	"gorm.io/gorm"
)

type Fornecedor struct {
	gorm.Model
	Nome     string `json:"nome" gorm:"column:nome;type:text"`
	Cnpj     string `json:"cnpj" gorm:"column:cnpj;type:text;unique"`
	Email    string `json:"email" gorm:"column:email;type:text;unique"`
	Endereco string `json:"endereco" gorm:"column:endereco;type:text"`
	Telefone string `json:"telefone" gorm:"column:telefone;unique;type:text;"`
}

func (f Fornecedor) TableName() string {
	return "fornecedor"
}

func (f *Fornecedor) Format() error {
	f.Nome = strings.TrimSpace(f.Nome)
	f.Cnpj = strings.TrimSpace(f.Cnpj)
	f.Email = strings.TrimSpace(f.Email)
	f.Endereco = strings.TrimSpace(f.Endereco)
	f.Telefone = strings.TrimSpace(f.Telefone)

	return f.validate()

}

func (f *Fornecedor) validate() error {
	switch {

	case !brdoc.IsCNPJ(f.Cnpj):
		{
			return errors.New("insira um CNPJ Valido")

		}
	case checkmail.ValidateFormat(f.Email) != nil:
		{
			return errors.New("insira um Email Valido")
		}

	case f.Endereco == "":
		{
			{
				return errors.New("insira um endereco Valido")
			}
		}
	case f.Telefone == "":
		{
			{
				return errors.New("insira um telefone Valido")
			}
		}
	default:
		return nil
	}
}
