package handler

import "../product"

type Handler struct {
	productFront product.Front
}

func NewHandler(pf product.Front) *Handler {
	return &Handler{
		productFront: pf,
	}
}
