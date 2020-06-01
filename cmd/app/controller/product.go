package controller

import (
	"../model"
	"github.com/jinzhu/gorm"
)

// ProductController : Product controller defination
type ProductController struct {
	db *gorm.DB
}

// NewProductController : Product controller instantiation
func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		db: db,
	}
}

func (pc *ProductController) GetAllProducts() (*model.Products, error) {

	var err error
	var p []model.Product

	// Run query
	err = pc.db.LogMode(true).Debug().Find(&p).Error

	// Error check
	if err != nil {

		// Checking for empty record
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		// If not empty and there is a genuine error
		return nil, err
	}

	var pi model.Products
	pi.Items = &p
	return &pi, err
}

// GetProductByName : Search product by the provided name
func (pc *ProductController) GetProductByName(name string) (*model.Products, error) {
	var err error

	// prepare a product model holder
	var p []model.Product

	// Run search
	err = pc.db.LogMode(true).Debug().Where("Name = ?", name).Find(&p).Error

	// Error check
	if err != nil {

		// Checking for empty record
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		// If not empty and there is a genuine error
		return nil, err
	}

	// Multi product holder
	var pi model.Products

	pi.Items = &p
	return &pi, err
}

func (pc *ProductController) GetProductByID(id string) (*model.Product, error) {
	var err error
	var p model.Product

	// Run query
	err = pc.db.LogMode(true).Debug().Where("Id = ?", id).First(&p).Error

	// Error check
	if err != nil {

		// Checking for empty record
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		// If not empty and there is a genuine error
		return nil, err
	}

	return &p, err
}

func (pc *ProductController) AddProduct(m *model.Product) error {
	var err error

	// Run create query
	err = pc.db.LogMode(true).Debug().Create(m).Error

	// Error check
	if err != nil {

		// Checking for empty record
		if gorm.IsRecordNotFoundError(err) {
			return err
		}

		// If not empty and there is a genuine error
		return err
	}

	return err
}

func (pc *ProductController) UpdateProduct(id string, cm *model.Product) error {
	var err error
	m := new(model.Product)
	m.ID = id
	// Run update query
	err = pc.db.LogMode(true).Debug().Model(&m).Updates(cm).Error

	// Error check
	if err != nil {
		// Checking for empty record
		if gorm.IsRecordNotFoundError(err) {
			return err
		}
	}

	return err
}

func (pc *ProductController) DeleteProduct(cm *model.Product) error {
	var err error
	// Run delete query
	err = pc.db.LogMode(true).Debug().Delete(&cm).Error

	// Error check
	if err != nil {

		// Checking for empty record
		if gorm.IsRecordNotFoundError(err) {
			return err
		}

		// If not empty and there is a genuine error
		return err
	}

	return err
}
