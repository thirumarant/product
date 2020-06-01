package model

// Product is the basis for product
type Product struct {
	ID            string  `gorm:"column:Id;type:varchar;primary_key" json:"Id"`
	Name          string  `gorm:"column:Name;type:varchar" json:"Name"`
	Description   string  `gorm:"column:Description;type:varchar" json:"Description"`
	Price         float64 `gorm:"column:Price;type:decimal(6,2)" json:"Price"`
	DeliveryPrice float64 `gorm:"column:DeliveryPrice;type:decimal(6,2)" json:"DeliveryPrice"`
}

// Products hold a list of product
type Products struct {
	Items *[]Product
}
