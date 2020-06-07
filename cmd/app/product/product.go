package product

import "../model"

type Front interface {
	List() model.ProductList
	ListByName(name string) model.ProductList
	GetByID(id string) (*model.Product, error)
	CreateProduct(*model.Product) error
	UpdateProduct(*model.Product) error
	DeleteProduct(*model.Product) error

	ListOptions(id string) model.ProductOptionList
	CreateOption(id string) error
	GetSpecificOption(id string, optionId string) (*model.Product, error)
	DeleteSpecificOption(id string, optionId string) error
	UpdateSpecificOption(id string, optionId string) error
}
