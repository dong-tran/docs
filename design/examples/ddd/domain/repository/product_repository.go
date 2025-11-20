package repository

import "github.com/dong-tran/docs/ddd-example/domain/model"

// ProductRepository defines the contract for product persistence
type ProductRepository interface {
	Save(product *model.Product) error
	FindByID(id model.ProductID) (*model.Product, error)
	FindAll() ([]*model.Product, error)
	Delete(id model.ProductID) error
}
