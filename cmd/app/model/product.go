package model

// Product model is the basis for product
type Product struct {
	ID            string          `gorm:"column:Id;type:varchar;primary_key" json:"Id" query:"id"`
	Name          string          `gorm:"column:Name;type:varchar" json:"Name" query:"Name"`
	Description   string          `gorm:"column:Description;type:varchar" json:"Description" query:"Description"`
	Price         float64         `gorm:"column:Price;type:decimal(6,2)" json:"Price" query:"Price"`
	DeliveryPrice float64         `gorm:"column:DeliveryPrice;type:decimal(6,2)" json:"DeliveryPrice" query:"DeliveryPrice"`
	ProductOption []ProductOption `gorm:"foreignkey:ProductId; association_foreignkey:Id" json:"-"`
}

// Product list holds an array of product models
type ProductList struct {
	Items []Product `json:"Items"`
}
