package model

// Product is the basis for product
type ProductOption struct {
	ID          string `gorm:"column:Id;type:varchar;primary_key" json:"Id" query:"id"`
	ProductID   string `gorm:"column:ProductId;type:varchar" json:"ProductId" query:"ProductId"`
	Name        string `gorm:"column:Name;type:varchar" json:"Name" query:"Name"`
	Description string `gorm:"column:Description;type:varchar" json:"Description" query:"Description"`
}

type ProductOptionList struct {
	Items *[]ProductOption `json:"Items"`
}
