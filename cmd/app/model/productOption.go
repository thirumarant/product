package model

// Product is the basis for product
type ProductOption struct {
	ID          string `gorm:"column:Id;type:varchar;primary_key" json:"Id"`
	ProductID   string `gorm:"column:ProductId;type:varchar" json:"ProductId"`
	Name        string `gorm:"column:Name;type:varchar" json:"Name"`
	Description string `gorm:"column:Description;type:varchar" json:"Description"`
}

type ProductOptionList struct {
	Items *[]ProductOption `json:"Items"`
}
