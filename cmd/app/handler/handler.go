package handler

import "../product"

// Main service handler container to hold interfaces
// To backend logic
type Handler struct {
	productFront product.Front
}

// Constructor for handler
func NewHandler(pf product.Front) *Handler {
	return &Handler{
		productFront: pf,
	}
}
