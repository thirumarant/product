package handler

import (
	"github.com/labstack/echo"
)

// Register : Merely a route registry to define and map routes against specific handler for processing
func (ph *ProductHandler) Register(v1 *echo.Group) {

	// `GET /products` - gets all products.
	// `GET /products?name={name}` - finds all products matching the specified name.
	v1.GET("", ph.Get)

	// `GET /products/{id}` - gets the product that matches the specified ID - ID is a GUID.
	v1.GET("/:id", ph.GetByID)

	// `POST /products` - creates a new product.
	v1.POST("", ph.Add)

	// `PUT /products/{id}` - updates a product.
	v1.PUT("/:id", ph.Update)

	// `DELETE /products/{id}` - deletes a product and its options.
	v1.DELETE("/:id", ph.Delete)

	// `GET /products/{id}/options` - finds all options for a specified product.
	// v1.GET("/:id/options", ph.FindAllOptionByProductID)
	v1.GET("/:id/options", ph.GetOptions)

	// `GET /products/{id}/options/{optionId}` - finds the specified product option for the specified product.
	v1.GET("/:id/options/:optionId", ph.GetAnOption)

	// `POST /products/{id}/options` - adds a new product option to the specified product.
	v1.POST("/:id/options", ph.AddAnOption)

	// `PUT /products/{id}/options/{optionId}` - updates the specified product option.
	v1.PUT("/:id/options/:optionId", ph.UpdateAnOption)

	//`DELETE /products/{id}/options/{optionId}` - deletes the specified product option.
	v1.DELETE("/:id/options/:optionId", ph.DeleteAnOption)
}
