// Package controller hold all controllers for
// this application
package controller

import (
	"../model"
	//"fmt"
	"github.com/jinzhu/gorm"
)

// ProductController is the holder of all product controller fields
type ProductController struct {
	Master
	db *gorm.DB
}

// NewProductController : Product controller instantiation
func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		db: db.LogMode(true).Debug(),
	}
}

func (pc *ProductController) GetAllProducts() (*model.ProductList, error) {
	return pc.processMultiRows()
}

// GetProductByName : Search product by the provided name
func (pc *ProductController) GetProductByName(name string) (*model.ProductList, error) {

	// Run search
	pc.db = pc.db.Where("Name = ?", name)

	if pc.db.Error == nil {
		return pc.processMultiRows()
	}

	return nil, pc.db.Error
}

func (pc *ProductController) GetProductByID(id string) (*model.Product, error) {
	var err error
	var p model.Product

	// Run query
	err = pc.db.Find(&p, id).Error

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
	// Run create query
	pc.db.Create(m)
	return pc.opsRecordCheck()

}

func (pc *ProductController) UpdateProduct(id string, cm *model.Product) error {
	m := new(model.Product)
	m.ID = id
	// Run update query
	pc.db.Model(&m).Updates(cm)
	return pc.opsRecordCheck()
}

func (pc *ProductController) DeleteProduct(cm *model.Product) error {
	// Run delete query
	pc.db.Delete(&cm)
	return pc.opsRecordCheck()
}

func (pc *ProductController) GetAllProductOption(pm *model.ProductOption) (*model.ProductOptionList, error) {
	var err error
	var po []model.ProductOption
	var poi model.ProductOptionList

	// Run query
	pc.db.Table("ProductOptions").Where("ProductId = ?", &pm.ProductID).Find(&po)

	// If there are no issues
	if pc.db.Error == nil {
		// TODO: Bug with the function below which seems to always return true
		// TODO: Circumventing the result at the handler to gracefully display empty results
		if !pc.db.RecordNotFound() {
			// package up results and return
			poi.Items = &po
			err = pc.db.Error
		}
	} else {
		err = pc.db.Error
	}

	// flush the db
	pc.db = pc.db.New()
	return &poi, err
}

func (pc *ProductController) opsRecordCheck() error {
	// Error check
	if pc.db.Error != nil {

		// Checking for empty record
		if gorm.IsRecordNotFoundError(pc.db.Error) {
			return pc.db.Error
		}

		// If not empty and there is a genuine error
		return pc.db.Error
	}

	return pc.db.Error
}

func (pc *ProductController) processMultiRows() (*model.ProductList, error) {

	var err error
	var p []model.Product
	var pi model.ProductList

	// Run query
	pc.db.Find(&p)

	// If there are no issues
	if pc.db.Error == nil {
		// TODO: Bug with the function below which seems to always return true
		// TODO: Circumventing the result at the handler to gracefully display empty results
		if !pc.db.RecordNotFound() {
			// package up results and return
			pi.Items = &p
			err = pc.db.Error
		}
	} else {
		err = pc.db.Error
	}

	// flush the db
	pc.db = pc.db.New()
	return &pi, err
}
