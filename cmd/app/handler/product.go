package handler

import (
	"net/http"

	"../model"
	"../utils"
	"github.com/labstack/echo"
)

func (h *Handler) Get(c echo.Context) (err error) {
	var productList model.ProductList

	name := c.QueryParam("name")

	if len(name) > 0 {
		productList, err = h.productFront.ListByName(name)
	} else {
		productList, err = h.productFront.List()
	}

	if len(productList.Items) == 0 {
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	// All good respond with results
	return c.JSONPretty(http.StatusOK, &productList, " ")
}

func (h *Handler) GetByID(c echo.Context) error {

	product, err := h.productFront.GetByID(c.Param("id"))

	if product == nil {
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	return c.JSONPretty(http.StatusOK, &product, " ")
}

// CreateProduct : Create a brand new product
func (h *Handler) Add(c echo.Context) (err error) {
	product := model.Product{}

	err = c.Bind(&product)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}
	err = h.productFront.CreateProduct(&product)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusCreated, map[string]interface{}{"result": "ok"}, " ")
}

// CreateProduct : Create a brand new product
func (h *Handler) Update(c echo.Context) (err error) {
	product := model.Product{ID: c.Param("id")}
	err = c.Bind(&product)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	err = h.productFront.UpdateProduct(&product)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"result": "ok"}, " ")
}

// CreateProduct : Create a brand new product
func (h *Handler) Delete(c echo.Context) (err error) {
	var product model.Product

	product.ID = c.Param("id")

	err = h.productFront.DeleteProduct(&product)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"result": "ok"}, " ")
}

func (h *Handler) GetOptions(c echo.Context) (err error) {

	var productOptionsList model.ProductOptionList

	productOptionsList, err = h.productFront.ListOptions(c.Param("id"))

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	return c.JSONPretty(http.StatusOK, &productOptionsList, " ")
}

// FindSpecificOptionByProductID : Retrieves the options of a product by the given product ID
func (h *Handler) GetAnOption(c echo.Context) error {

	productOption, err := h.productFront.GetSpecificOption(c.Param("id"), c.Param("optionId"))

	if productOption == nil {
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	return c.JSONPretty(http.StatusOK, &productOption, " ")
}

// AddOptionByProductID : Adds option for a product by the given product ID
func (h *Handler) AddAnOption(c echo.Context) (err error) {
	productOption := model.ProductOption{ProductID: c.Param("id")}
	err = c.Bind(&productOption)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	err = h.productFront.CreateOption(&productOption)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusCreated, map[string]interface{}{"result": "ok"}, " ")
}

// UpdateSpecificOptionByProductID : Updates a specific option of a product by the given product ID
func (h *Handler) UpdateAnOption(c echo.Context) (err error) {
	productOption := model.ProductOption{}
	err = c.Bind(&productOption)
	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	err = h.productFront.UpdateSpecificOption(c.Param("id"), c.Param("optionId"), &productOption)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"result": "ok"}, " ")
}

// DeleteSpecificOptionByProductID : Removes a specific option of a product by the given product ID
func (h *Handler) DeleteAnOption(c echo.Context) (err error) {

	err = h.productFront.DeleteSpecificOption(c.Param("id"), c.Param("optionId"))

	if err != nil {

		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"result": "ok"}, " ")
}
