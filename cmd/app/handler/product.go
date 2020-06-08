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

// Get product options retrieves the options of a product
// return error
// Router /products/{id}/options [get]
func (h *Handler) GetOptions(c echo.Context) (err error) {

	// Model
	var productOptionsList model.ProductOptionList

	// Run the controller function and hydrate the model
	productOptionsList, err = h.productFront.ListOptions(c.Param("id"))

	// check for processing errors
	if err != nil {

		// Return issues
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	// All good response with results
	return c.JSONPretty(http.StatusOK, &productOptionsList, " ")
}

// Get a specific product option retrieves the a particular option of a product
// return error
// Router /products/{id}/options/{optionId} [get]
func (h *Handler) GetAnOption(c echo.Context) error {

	// Run controller with filters to retrieve results and populate model
	productOption, err := h.productFront.GetSpecificOption(c.Param("id"), c.Param("optionId"))

	// If the model didn't get populated
	if productOption == nil {

		// No is found with specification 404
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	// Check for processing error as well
	if err != nil {

		// Notify about error
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	// All good response with results
	return c.JSONPretty(http.StatusOK, &productOption, " ")
}

// Add an option to the product add a specific option to a given product
// returns error
// Router /products/{id}/options [post]
func (h *Handler) AddAnOption(c echo.Context) (err error) {

	// Prepare a model with relevant product ID
	productOption := model.ProductOption{ProductID: c.Param("id")}

	// Bind incoming payload with model
	// TODO: This is where input validation needs to be introduced
	err = c.Bind(&productOption)

	// Check for binding errors
	if err != nil {

		// Response errors as a conflicts
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	// Inject model into controller to create
	err = h.productFront.CreateOption(&productOption)

	// Check for creation error
	if err != nil {

		// Return issue
		// TODO: Need to decide on the type of error response to be specific
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	// All good response
	return c.JSONPretty(http.StatusCreated, map[string]interface{}{"result": "ok"}, " ")
}

// Update a product option modifies an existing specific option of a given product
// return error
// Router /products/{id}/options/{optionId} [put]
func (h *Handler) UpdateAnOption(c echo.Context) (err error) {

	// Prepare a model
	productOption := model.ProductOption{}

	// Bind incoming payload with model
	err = c.Bind(&productOption)

	// Check for binding errors
	if err != nil {

		// Return issues
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	// Run controller function to update using filters
	err = h.productFront.UpdateSpecificOption(c.Param("id"), c.Param("optionId"), &productOption)

	// Check for controller processing errors
	if err != nil {

		// Return issues
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	// All good response
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"result": "ok"}, " ")
}

// Delete an option of a product removes a specific option of a given product
// return error
// Router /products/{id}/options/{optionId} [delete]
func (h *Handler) DeleteAnOption(c echo.Context) (err error) {

	// Run controller function with filters
	err = h.productFront.DeleteSpecificOption(c.Param("id"), c.Param("optionId"))

	// Check for processing error
	if err != nil {

		// Return issues
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	// All good response
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"result": "ok"}, " ")
}
