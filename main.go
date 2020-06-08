package main

import (
	"./cmd/app/controller"
	"./cmd/app/handler"
	"./cmd/app/router"
	"./cmd/app/storage"
)

func main() {

	// Instantiate the HTTP Framework to manage service
	r := router.New()

	// Group the service name
	v1 := r.Group("/products")

	// Instantiate a new storage to be injected
	db := storage.New()

	// Instantiate the service controller
	c := controller.NewProductController(db)

	// Instantiate the web handler and inject necessary components
	handler.NewHandler(c).Register(v1)

	// Start the web server
	r.Logger.Fatal(r.Start("127.0.0.1:8080"))
}
