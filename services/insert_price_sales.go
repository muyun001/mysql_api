package services

import (
	"mysql_api/databases"
	"mysql_api/structs"
	"mysql_api/structs/models"
	"time"
)

func InsertPriceSales(info *structs.ProductPriceSales) error {
	priceSales := &models.PriceSales{
		ProductCode: info.ProductCode,
		Price:       info.Price,
		Sales:       info.Sales,
		Date:        info.Date,
		CreatedAt:   time.Time{},
	}
	if err := databases.Db.Save(priceSales).Error; err != nil {
		return err
	}

	product := &models.Product{}
	databases.Db.Where(&models.Product{ProductCode: info.ProductCode}).Find(&product)

	err := databases.Db.Model(product).Updates(&models.Product{
		Status:     models.StatusSearchSucceed,
		TotalSales: product.TotalSales + info.Sales,
	}).Error

	return err
}
