package handler

import (
	"net/http"

	"../controller"
	"../model"
	"../utils"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// ProductHandler container field holder
type ProductHandler struct {
	pc *controller.ProductController
}

// NewProductHandler is the constructor to instantiate a product handler
func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{
		pc: controller.NewProductController(db),
	}
}

// GetAllProducts : Retrieve all products and can handle a name filter
func (ph *ProductHandler) GetAllProducts(c echo.Context) (err error) {

	// Get a pointer to the product model
	var pm *model.ProductList

	// Check for the presence of the name query string
	n := c.QueryParam("name")

	// If name exists url query
	// filter using name else unfiltered result
	if len(n) > 0 {
		pm, err = ph.pc.GetProductByName(n)
	} else {
		pm, err = ph.pc.GetAllProducts()
	}

	// Check for controller issues and throw an internal error
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	// If results are empty throw a not found status
	if pm == nil {
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	// All good respond with results
	return c.JSONPretty(http.StatusOK, pm, " ")
}

// GetProductByID : Retrieves a specific product by id
func (ph *ProductHandler) GetProductByID(c echo.Context) error {
	var err error

	pm, err := ph.pc.GetProductByID(c.Param("id"))

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	if pm == nil {
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	return c.JSONPretty(http.StatusOK, pm, " ")
}

// CreateProduct : Create a brand new product
func (ph *ProductHandler) CreateProduct(c echo.Context) error {
	var err error
	m := new(model.Product)
	err = c.Bind(m)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	err = ph.pc.AddProduct(m)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSONPretty(http.StatusCreated, map[string]interface{}{"result": "ok"}, " ")
}

// UpdateProductByID : Updates the product details by the provided product ID
func (ph *ProductHandler) UpdateProductByID(c echo.Context) error {
	var err error

	// Get a new model
	m := new(model.Product)

	// Merge the body with the new struct
	err = c.Bind(m)

	// Check for bind issue
	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	// Get the controller to run update
	err = ph.pc.UpdateProduct(c.Param("id"), m)

	// Check for controller issues
	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

// DeleteProductByID : Removes a product using the provided product ID
func (ph *ProductHandler) DeleteProductByID(c echo.Context) error {
	var err error

	// Get a new model
	m := new(model.Product)

	// Grab the relevant product id
	m.ID = c.Param("id")

	// get the controller to run the delete
	err = ph.pc.DeleteProduct(m)

	// Check for issues
	if err != nil {
		return c.JSON(http.StatusConflict, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

// FindAllOptionByProductID : Retrieves all the options of a product by the given product ID
func (ph *ProductHandler) FindAllOptionByProductID(c echo.Context) error {
	var err error

	po := new(model.ProductOption)

	po.ProductID = c.Param("id")

	pl, err := ph.pc.GetAllProductOption(po)

	// Check for controller issues and throw an internal error
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	// If results are empty throw a not found status
	if po == nil {
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	// All good respond with results
	return c.JSONPretty(http.StatusOK, pl, " ")
}

// FindSpecificOptionByProductID : Retrieves the options of a product by the given product ID
func (ph *ProductHandler) FindSpecificOptionByProductID(c echo.Context) error {
	var err error

	po := new(model.ProductOption)

	po.ProductID = c.Param("id")
	po.ID = c.Param("optionId")

	pl, err := ph.pc.GetAllProductOption(po)

	// Check for controller issues and throw an internal error
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	// If results are empty throw a not found status
	if po == nil {
		return c.JSONPretty(http.StatusNotFound, utils.NotFound(), " ")
	}

	// All good respond with results
	return c.JSONPretty(http.StatusOK, pl, " ")
}

// AddOptionByProductID : Adds option for a product by the given product ID
func (ph *ProductHandler) AddOptionByProductID(c echo.Context) error {
	var err error

	return err
}

// UpdateSpecificOptionByProductID : Updates a specific option of a product by the given product ID
func (ph *ProductHandler) UpdateSpecificOptionByProductID(c echo.Context) error {
	var err error

	return err
}

// DeleteSpecificOptionByProductID : Removes a specific option of a product by the given product ID
func (ph *ProductHandler) DeleteSpecificOptionByProductID(c echo.Context) error {
	var err error

	return err
}
