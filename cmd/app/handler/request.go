package handler

import (
	"../model"
	"../utils"
	"errors"
	"github.com/labstack/echo"
	"net/http"
	"regexp"
	"strconv"
)

// Structs for mapping incoming json payload
type ProductRequestPayload struct {
	ID            string  `json:"Id"`
	Name          string  `json:"Name"`
	Description   string  `json:"Description"`
	Price         float64 `json:"Price"`
	DeliveryPrice float64 `json:"DeliveryPrice"`
}

type ProductOptionRequestPayload struct {
	ID          string `json:"Id"`
	ProductID   string `json:"ProductId"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// ValidateProductPayload check for the data validity of the json payload values
// returns error
func (h *Handler) ValidateProductPayload(c echo.Context, model *model.Product) error {
	var err error
	var msg string
	var rp ProductRequestPayload
	err = c.Bind(&rp)

	// Check for binding error
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	// map to model
	model.ID = rp.ID
	model.Name = rp.Name
	model.Description = rp.Description
	model.Price = rp.Price
	model.DeliveryPrice = rp.DeliveryPrice

	if len(rp.ID) > 0 {
		msg = "Id is system generated, please do not supply"
	}

	if len(rp.Name) < 1 || len(rp.Name) > 17 {
		msg = "Product Name length should be between 1-17 characters"
	}

	if len(rp.Description) > 35 {
		msg = "Description may be between 1-36 characters"
	}
	if err = h.validatePrices(rp.Price, "price"); err != nil {
		msg = err.Error()
	}
	if err = h.validatePrices(rp.DeliveryPrice, "delivery price"); err != nil {
		msg = err.Error()
	}

	if len(msg) > 0 {
		return errors.New(msg)
	}

	return nil
}

// ValidateProductOptionPayload check for the data validity of the json payload values
// returns error
func (h *Handler) ValidateProductOptionPayload(c echo.Context, model *model.ProductOption) error {
	var msg string
	var rpo ProductOptionRequestPayload
	err := c.Bind(&rpo)

	// Check for binding error
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, utils.NewError(err), " ")
	}

	// Fix framework assignment issue
	// Todo: Refactor the routes
	rpo.ID = ""

	// Pass the param product ID
	rpo.ProductID = model.ProductID

	// Exchange data with model
	model.Name = rpo.Name
	model.Description = rpo.Description

	if len(rpo.ID) > 0 {
		msg = "Id is system generated, please do not supply"
	}
	if len(rpo.ProductID) == 0 {
		msg = "Please provide a product id"
	}
	if !utils.IsValidUUID(rpo.ProductID) {
		msg = "Product id id not valid UUID"
	}
	if len(rpo.Name) < 1 || len(rpo.Name) > 17 {
		msg = "Product option Name length should be between 1-17 characters"
	}
	if len(rpo.Description) > 35 {
		msg = "Description may be between 1-36 characters"
	}

	if len(msg) > 0 {
		return errors.New(msg)
	}

	return nil
}

// validatePrice check for the data validity of the provided prices
// returns boolean
func (h *Handler) validatePrices(price float64, key string) error {
	priceStr := strconv.FormatFloat(price, 'f', -1, 64)
	if len(priceStr) > 0 {
		priceMatch := regexp.MustCompile(`^[0-9]+([.][0-9]{1,2})?$`)
		if !priceMatch.MatchString(priceStr) {
			return errors.New(key + " is not a valid decimal number")
		}
	}
	return nil
}
