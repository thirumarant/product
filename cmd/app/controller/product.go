package controller

import (
	"../model"
	"../utils"
	"github.com/jinzhu/gorm"
)

// This is the product controller section
// All business logic necessary for the CRUD functionality of the product
// sits in here
// This layer applies the necessary processing and orchestration needed for the data for storage
// and relevant enrichment for presentation and experience for the calling layer

// Controller field holder
type ProductController struct {
	db *gorm.DB
}

// Constructor returning an instance of the controller which carries the injected DB
func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		db: db,
	}
}

func (pc *ProductController) List() (model.ProductList, error) {
	var products []model.Product

	pc.db.Find(&products)

	return model.ProductList{Items: products}, pc.db.Error
}

func (pc *ProductController) ListByName(name string) (model.ProductList, error) {
	var products []model.Product

	pc.db.Where(&model.Product{Name: name}).Find(&products)

	return model.ProductList{Items: products}, pc.db.Error
}

func (pc *ProductController) GetByID(id string) (*model.Product, error) {
	var product model.Product

	err := pc.db.Where("Id = ?", id).Find(&product).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, err
	}

	return &product, nil
}

func (pc *ProductController) CreateProduct(product *model.Product) error {
	product.ID = utils.GenerateUUID()
	return pc.db.Create(&product).Error
}

func (pc *ProductController) UpdateProduct(product *model.Product) error {
	if pc.db.Model(&model.Product{ID: product.ID}).Updates(&product).RowsAffected == 0 && &pc.db.Error == nil {
		return gorm.ErrRecordNotFound
	}

	return pc.db.Error
}

func (pc *ProductController) DeleteProduct(product *model.Product) error {
	if pc.db.Delete(&product).RowsAffected == 0 && pc.db.Error == nil {
		return gorm.ErrRecordNotFound
	}

	return pc.db.Error
}

func (pc *ProductController) ListOptions(id string) (model.ProductOptionList, error) {
	var productOptions []model.ProductOption

	pc.db.Table("ProductOptions").
		Where("ProductId = ?", id).Find(&productOptions)

	return model.ProductOptionList{Items: productOptions}, pc.db.Error
}

func (pc *ProductController) CreateOption(productOption *model.ProductOption) error {
	productOption.ID = utils.GenerateUUID()
	return pc.db.Table("ProductOptions").Create(&productOption).Error
}

func (pc *ProductController) GetSpecificOption(id string, optionId string) (*model.ProductOption, error) {
	var productOption model.ProductOption

	if pc.db.Table("ProductOptions").
		Where("ProductId = ? AND Id = ?", id, optionId).
		Find(&productOption).RowsAffected == 0 {
		if pc.db.Error != nil {
			return nil, pc.db.Error
		}
		return nil, nil
	}

	return &productOption, pc.db.Error
}

func (pc *ProductController) UpdateSpecificOption(id string, optionId string, po *model.ProductOption) error {
	if pc.db.Table("ProductOptions").
		Where("Id = ? AND ProductId = ?", optionId, id).
		Model(model.ProductOption{}).
		Omit("Id").
		Updates(&po).
		RowsAffected == 0 && pc.db.Error == nil {
		return gorm.ErrRecordNotFound
	}

	return pc.db.Error
}

func (pc *ProductController) DeleteSpecificOption(id string, optionId string) error {
	if pc.db.Table("ProductOptions").
		Delete(&model.ProductOption{ProductID: id, ID: optionId}).
		RowsAffected == 0 && pc.db.Error == nil {
		return gorm.ErrRecordNotFound
	}

	return pc.db.Error
}
