package model

// Product option model is the basis for product options
type ProductOption struct {
	ID          string `gorm:"column:Id;type:varchar;primary_key" json:"Id" query:"id"`
	ProductID   string `gorm:"column:ProductId;type:varchar" json:"ProductId" query:"ProductId"`
	Name        string `gorm:"column:Name;type:varchar" json:"Name" query:"Name"`
	Description string `gorm:"column:Description;type:varchar" json:"Description" query:"Description"`
}

// Product option list holds an array of product option models
type ProductOptionList struct {
	Items []ProductOption `json:"Items"`
}
