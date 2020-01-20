package databases

import "mysql_api/structs/models"

func AutoMigrate() {
	Db.AutoMigrate(&models.Product{}, &models.PriceSales{})
}
