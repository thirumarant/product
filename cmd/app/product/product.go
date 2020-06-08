package product

import "../model"

// This is the interface which acts as an abstraction layer to the business logic
// it allows for the controlled exposure and presentation of functionality needed
// necessary to interact with the experience layer
type Front interface {
	// Core product functionality
	List() (model.ProductList, error)
	ListByName(name string) (model.ProductList, error)
	GetByID(id string) (*model.Product, error)
	CreateProduct(*model.Product) error
	UpdateProduct(*model.Product) error
	DeleteProduct(*model.Product) error

	// Necessary product options functionality
	ListOptions(id string) (model.ProductOptionList, error)
	CreateOption(*model.ProductOption) error
	GetSpecificOption(id string, optionId string) (*model.ProductOption, error)
	UpdateSpecificOption(id string, optionId string, po *model.ProductOption) error
	DeleteSpecificOption(id string, optionId string) error
}
