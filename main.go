package main

import (
	"./cmd/app/controller"
	"./cmd/app/handler"
	"./cmd/app/router"
	"./cmd/app/storage"
)

func main() {
	r := router.New()
	v1 := r.Group("/products")
	db := storage.New()
	c := controller.NewProductController(db)
	handler.NewHandler(c).Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
