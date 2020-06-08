package handler

import (
	"net/http"

	"../model"
	"../utils"
	"github.com/labstack/echo"
)

// Product specific handler specification
// All product related handlers are defined and
// managed in here

// Get product handler retrieves all products
// returns an error
// Router /products or /products?name={} [get]
func (h *Handler) Get(c echo.Context) (err error) {

	// Prepare model
	var productList model.ProductList

	// TODO: Check if the query parameter exists at all

	// Get query parameter name
	name := c.QueryParam("name")

	// Check if a name was given
	if len(name) > 0 {
		// List products by name
		productList, err = h.productFront.ListByName(name)
	} else {
		// List all products
		productList, err = h.productFront.List()
	}

	// Check if any results came back
	if len(productList.Items) == 0 {

		// 404 nothing found
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	// Check if any error got thrown during processing
	if err != nil {

		// Format error for response
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	// All good respond with results
	return c.JSONPretty(http.StatusOK, &productList, " ")
}

// Get product by the given product ID
// return error
// Router /products/{id} [get]
func (h *Handler) GetByID(c echo.Context) error {

	// Run controller to pull results
	product, err := h.productFront.GetByID(c.Param("id"))

	// Check if anything came back
	if product == nil {

		// If empty response 404
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	// Check for processing error
	if err != nil {

		// Format error for response
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	// All good respond with results
	return c.JSONPretty(http.StatusOK, &product, " ")
}

// Add product creates a brand new product
// returns error
// Router /products [post]
func (h *Handler) Add(c echo.Context) (err error) {

	// Get model
	product := model.Product{}

	// bind json payload
	err = c.Bind(&product)

	// Check for binding issues to bail out
	if err != nil {
		// Return a conflict status
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	// Proceed to create product with controller
	err = h.productFront.CreateProduct(&product)

	// Check for processing errors
	if err != nil {

		// Return formatted response
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	// All good respond
	return c.JSONPretty(http.StatusCreated, map[string]interface{}{"result": "ok"}, " ")
}

// Update a product updates an existing product
// returns error
// Router /products/{id} [put]
func (h *Handler) Update(c echo.Context) (err error) {

	// Instantiate a model with incoming product ID
	product := model.Product{ID: c.Param("id")}

	// Bind the payload to the model
	err = c.Bind(&product)

	// Check for binding error
	if err != nil {

		// Response with conflict stating the issue
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	// Run the controller for update
	err = h.productFront.UpdateProduct(&product)

	// Check for processing error
	if err != nil {

		// Return conflicts
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	// All good respond
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"result": "ok"}, " ")
}

// Delete a product removes a product from storage
// return error
// Router /products/{id} [delete]
func (h *Handler) Delete(c echo.Context) (err error) {

	// Get model
	var product model.Product

	// Get incoming product id
	product.ID = c.Param("id")

	// Get controller to run delete
	err = h.productFront.DeleteProduct(&product)

	// Check for processing error
	if err != nil {

		// Response issue with correct code
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	// All good response
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
