package models

import "time"

// 商品表
type Product struct {
	ID               int       `gorm:"primary_key" json:"id"`
	ProductId        int       `gorm:"type:int" json:"product_id"`
	ProductCode      string    `gorm:"type:varchar(32);unique_index;index" json:"product_code"`
	Status           int       `gorm:"type:int;default:0;index" json:"status"`
	ProductLink      string    `gorm:"type:varchar(255)" json:"product_link"`
	ProductTitle     string    `gorm:"type:varchar(255)" json:"product_title"`
	ProductImg       string    `gorm:"type:varchar(255)" json:"product_img"`
	OnlineTime       string    `gorm:"type:varchar(32)" json:"online_time"`
	ProductCreatedAt string    `gorm:"type:varchar(32)" json:"product_created_at"`
	ProductUpdatedAt string    `gorm:"type:varchar(32)" json:"product_updated_at"`
	FirstType        string    `gorm:"type:varchar(32)" json:"first_type"`
	SecondType       string    `gorm:"type:varchar(32)" json:"second_type"`
	TotalSales       int       `gorm:"type:int" json:"total_sales"`
	ShopName         string    `gorm:"type:varchar(64)" json:"shop_name"`
	ShopCode         string    `gorm:"type:varchar(32)" json:"shop_code"`
	ShopLink         string    `gorm:"type:varchar(255)" json:"shop_link"`
	ShopLogo         string    `gorm:"type:varchar(255)" json:"shop_logo"`
	ShopCreatedAt    string    `gorm:"type:varchar(32)" json:"shop_created_at"`
	ShopUpdatedAt    string    `gorm:"type:varchar(32)" json:"shop_updated_at"`
	CreatedAt        time.Time `json:"created_at"`
}

// 价格销量表
type PriceSales struct {
	ID          int       `gorm:"primary_key" json:"id"`
	ProductCode string    `gorm:"type:varchar(32);index;unique_index:code_date" json:"product_code"`
	Price       float64   `gorm:"type:float" json:"price"`
	Sales       int       `gorm:"type:int" json:"sales"`
	Date        string    `gorm:"type:varchar(16);unique_index:code_date" json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	Product     Product
}
