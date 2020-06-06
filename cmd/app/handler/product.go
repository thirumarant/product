package handler

import (
	"net/http"

	"../model"
	"../utils"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// ProductHandler container field holder
type ProductHandler struct {
	db *gorm.DB
}

// NewProductHandler is the constructor to instantiate a product handler
func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{
		db: db,
	}
}

func (ph *ProductHandler) Get(c echo.Context) (err error) {
	ph.db = ph.db.New()
	products := new([]model.Product)

	name := c.QueryParam("name")

	if len(name) > 0 {
		ph.db = ph.db.Where("name = ?", name).Find(&products)
	} else {
		ph.db = ph.db.Find(&products)
	}

	// Check for controller issues and throw an internal error
	if ph.db.RowsAffected < 1 {
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	if ph.db.Error != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(ph.db.Error), " ")
	}

	// All good respond with results
	return c.JSONPretty(http.StatusOK, &model.ProductList{Items: products}, " ")
}

func (ph *ProductHandler) GetByID(c echo.Context) (err error) {
	ph.db.New()
	product := new(model.Product)

	// Run query
	ph.db = ph.db.Where("Id = ?", c.Param("id")).Find(&product)

	if ph.db.Error != nil {
		if ph.db.RowsAffected < 1 {
			return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
		}

		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	return c.JSONPretty(http.StatusOK, &product, " ")
}

// CreateProduct : Create a brand new product
func (ph *ProductHandler) Add(c echo.Context) (err error) {
	product := model.Product{}
	err = c.Bind(&product)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	err = ph.db.Create(&product).Error

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusCreated, map[string]interface{}{"result": "ok"}, " ")
}

// CreateProduct : Create a brand new product
func (ph *ProductHandler) Update(c echo.Context) (err error) {
	product := new(model.Product)
	err = c.Bind(product)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	err = ph.db.Model(&model.Product{ID: c.Param("id")}).Updates(&product).Error

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"result": "ok"}, " ")
}

// CreateProduct : Create a brand new product
func (ph *ProductHandler) Delete(c echo.Context) (err error) {
	product := model.Product{ID: c.Param("id")}

	err = ph.db.Delete(&product).Error

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"result": "ok"}, " ")
}

func (ph *ProductHandler) GetOptions(c echo.Context) (err error) {

	var productOptions []model.ProductOption

	productId := c.Param("id")

	// Run query
	ph.db.Table("ProductOptions").Where("ProductId = ?", productId).Find(&productOptions)

	if ph.db.Error != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	if &productOptions == nil {
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	return c.JSONPretty(http.StatusOK, model.ProductOptionList{Items: &productOptions}, " ")
}

// FindSpecificOptionByProductID : Retrieves the options of a product by the given product ID
func (ph *ProductHandler) GetAnOption(c echo.Context) (err error) {
	var productOption model.ProductOption

	productId := c.Param("id")
	optionId := c.Param("optionId")

	query := "Id = ? AND ProductId = ?"
	// Run query
	ph.db = ph.db.Table("ProductOptions").Where(query, optionId, productId).Find(&productOption)

	if ph.db.Error != nil {
		if ph.db.RowsAffected < 1 {
			return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
		}

		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(ph.db.Error), " ")
	}

	return c.JSONPretty(http.StatusOK, &productOption, " ")
}

// AddOptionByProductID : Adds option for a product by the given product ID
func (ph *ProductHandler) AddAnOption(c echo.Context) (err error) {
	productOption := model.ProductOption{ProductID: c.Param("id")}
	err = c.Bind(&productOption)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	err = ph.db.Table("ProductOptions").Create(&productOption).Error

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusCreated, map[string]interface{}{"result": "ok"}, " ")
}

// UpdateSpecificOptionByProductID : Updates a specific option of a product by the given product ID
func (ph *ProductHandler) UpdateAnOption(c echo.Context) (err error) {

	productOption := model.ProductOption{}

	err = c.Bind(productOption)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	query := "Id = ? AND ProductId = ?"

	ph.db.Model(&model.ProductOption{}).Where(query, c.Param("optionId"), c.Param("id")).Updates(&productOption)

	err = ph.db.Error

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"result": "ok"}, " ")
}

// DeleteSpecificOptionByProductID : Removes a specific option of a product by the given product ID
func (ph *ProductHandler) DeleteAnOption(c echo.Context) (err error) {
	productOption := model.ProductOption{}
	productOption.ProductID = c.Param("id")
	productOption.ID = c.Param("optionId")

	ph.db = ph.db.Table("ProductOptions").Delete(&productOption)

	if ph.db.Error != nil {

		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	if ph.db.RowsAffected < 1 {
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"result": "ok"}, " ")
}
