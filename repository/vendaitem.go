package repository

import (
	"api/db"
	"api/models"
	"errors"

	"gorm.io/gorm"
)

type ItemVendaRepo[T any] interface {
	Create(T) error
	List() ([]T, error)
	Find(id uint) error
	Update(T) error
	Delete(id uint) error
}

type itemVenda struct {
	db *gorm.DB
}

func (i *itemVenda) Create(iv models.ItemVenda) error {
	vendaRepo, err := NewVendaRepo()
	if err != nil {
		return err
	}
	venda, err := vendaRepo.Find(iv.VendaID)
	if err != nil {
		return err
	}

	if venda.Concluida {
		return errors.New("nao eh possivel adicionar itens a uma venda ja concluida")
	}
	return i.db.Create(iv).Error
}

func (i *itemVenda) Delete(id uint) error {
	var item models.ItemVenda
	item.ID = id
	return i.db.Delete(item).Error
}

func (i *itemVenda) Find(id uint) error {
	panic("unimplemented")
}

func (i *itemVenda) List() ([]models.ItemVenda, error) {
	var itens []models.ItemVenda
	if err := i.db.Find(itens).Error; err != nil {
		return nil, err
	}
	return itens, nil
}

func (i *itemVenda) Update(v models.ItemVenda) error {
	if err := i.db.Where("id=?", v.ID).Updates(v).Error; err != nil {
		return err
	}
	return nil
}

func NewItemVendaRepo() (ItemVendaRepo[models.ItemVenda], error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	return &itemVenda{db: db.Model(&models.ItemVenda{})}, nil
}
