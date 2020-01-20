package structs

// 接收商品信息的请求
type ProductInfo struct {
	ProductId        int     `gorm:"type:int" json:"product_id"`
	ProductCode      string  `gorm:"type:varchar(64);index" json:"product_code"`
	Status           int     `gorm:"type:int;default:0;index" json:"status"`
	ProductLink      string  `gorm:"type:varchar(255)" json:"product_link"`
	ProductTitle     string  `gorm:"type:varchar(255)" json:"product_title"`
	ProductImg       string  `gorm:"type:varchar(255)" json:"product_img"`
	OnlineTime       string  `gorm:"type:varchar(32)" json:"online_time"`
	Price            float64 `gorm:"type:float" json:"price"`
	TodaySales       int     `gorm:"type:int" json:"today_sales"`
	TotalSales       int     `gorm:"type:int" json:"total_sales"`
	ProductCreatedAt string  `gorm:"type:varchar(32)" json:"product_created_at"`
	ProductUpdatedAt string  `gorm:"type:varchar(32)" json:"product_updated_at"`
	FirstType        string  `gorm:"type:varchar(64)" json:"first_type"`
	SecondType       string  `gorm:"type:varchar(64)" json:"second_type"`
	ShopName         string  `gorm:"type:varchar(255)" json:"shop_name"`
	ShopCode         string  `gorm:"type:varchar(32)" json:"shop_code"`
	ShopLink         string  `gorm:"type:varchar(255)" json:"shop_link"`
	ShopLogo         string  `gorm:"type:varchar(255)" json:"shop_logo"`
	ShopCreatedAt    string  `gorm:"type:varchar(32)" json:"shop_created_at"`
	ShopUpdatedAt    string  `gorm:"type:varchar(32)" json:"shop_updated_at"`
	Date             string  `gorm:"type:varchar(16)" json:"date"`
}

// 用于接收每天插入的商品价格和销量的请求
type ProductPriceSales struct {
	ProductCode string  `gorm:"type:varchar(64);index" json:"product_code"`
	Price       float64 `gorm:"type:float" json:"price"`
	Sales       int     `gorm:"type:int" json:"sales"`
	Date        string  `gorm:"type:varchar(16)" json:"date"`
}

// 接收更新状态的请求
type UpdateStatusRequest struct {
	ProductCode string `json:"product_code"`
	Status      int    `json:"status"`
}
