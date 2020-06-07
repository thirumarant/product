package controller

import (
	"../model"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/poly1305"
)

type ProductController struct {
	db *gorm.DB
}

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

	if err := pc.db.Where("Id = ?", id).Find(&product).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, err
	}

	return &product, nil
}

func (pc *ProductController) CreateProduct(product *model.Product) error {
	return pc.db.Create(&product).Error
}

func (pc *ProductController) UpdateProduct(product *model.Product) error {
	return pc.db.Model(&model.Product{ID: product.ID}).Updates(&product).Error
}

func (pc *ProductController) DeleteProduct(product *model.Product) error {
	return pc.db.Delete(&product).Error
}

func (pc *ProductController) ListOptions(id string) (model.ProductOptionList, error) {
	var productOptions []model.ProductOption

	pc.db.Table("ProductOptions").
		Where(&model.ProductOption{ProductID: id}).Find(&productOptions)

	return model.ProductOptionList{Items: &productOptions}, pc.db.Error
}

func (pc *ProductController) CreateOption(id string) error {
	var productOption model.ProductOption

	return pc.db.Table("ProductOptions").Create(&productOption).Error
}

func (pc *ProductController) GetSpecificOption(id string, optionId string) (*model.ProductOption, error) {
	var productOption model.ProductOption

	pc.db.Table("ProductOptions").
		Where(&model.ProductOption{ProductID: id, ID: optionId}).Find(&productOption)

	return &productOption, pc.db.Error
}

func (pc *ProductController) UpdateSpecificOption(po model.ProductOption) error {
	return pc.db.Where(&model.ProductOption{ProductID: po.ProductID, ID: po.ID}).
		Model(&model.ProductOption{}).
		Updates(&po).Error
}

func (pc *ProductController) DeleteSpecificOption(id string, optionId string) error {
	return pc.db.Table("ProductOptions").Delete(&model.ProductOption{ProductID: id, ID: optionId}).Error
}
