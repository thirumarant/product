package product

import "../model"

type Front interface {
	List() (model.ProductList, error)
	ListByName(name string) (model.ProductList, error)
	GetByID(id string) (*model.Product, error)
	CreateProduct(*model.Product) error
	UpdateProduct(*model.Product) error
	DeleteProduct(*model.Product) error

	ListOptions(id string) (model.ProductOptionList, error)
	CreateOption(*model.ProductOption) error
	GetSpecificOption(id string, optionId string) (*model.ProductOption, error)
	UpdateSpecificOption(id string, optionId string, po *model.ProductOption) error
	DeleteSpecificOption(id string, optionId string) error
}
