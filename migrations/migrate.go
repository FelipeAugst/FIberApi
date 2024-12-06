package migrations

import (
	"api/db"
	"api/models"
)

func MigrateTables() error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	return db.AutoMigrate(&models.Peca{}, &models.Cliente{}, &models.Fornecedor{}, &models.Venda{}, &models.PedidoDeVenda{})

}
