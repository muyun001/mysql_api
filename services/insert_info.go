package services

import (
	"mysql_api/databases"
	"mysql_api/structs"
	"mysql_api/structs/models"
	"time"
)

// 保存数据
func InsertInfo(info *structs.ProductInfo) error {
	product := &models.Product{
		ProductId:        info.ProductId,
		ProductCode:      info.ProductCode,
		Status:           info.Status,
		ProductLink:      info.ProductLink,
		ProductTitle:     info.ProductTitle,
		ProductImg:       info.ProductImg,
		OnlineTime:       info.OnlineTime,
		ProductCreatedAt: info.ProductCreatedAt,
		ProductUpdatedAt: info.ProductUpdatedAt,
		FirstType:        info.FirstType,
		SecondType:       info.SecondType,
		TotalSales:       info.TotalSales,
		ShopName:         info.ShopName,
		ShopCode:         info.ShopCode,
		ShopLink:         info.ShopLink,
		ShopLogo:         info.ShopLogo,
		ShopCreatedAt:    info.ShopCreatedAt,
		ShopUpdatedAt:    info.ShopUpdatedAt,
		CreatedAt:        time.Time{},
	}
	priceSales := &models.PriceSales{
		ProductCode: info.ProductCode,
		Price:       info.Price,
		Sales:       info.TotalSales,
		Date:        info.Date,
		CreatedAt:   time.Time{},
		Product:     models.Product{},
	}

	err := databases.Db.Save(product).Error
	err = databases.Db.Save(priceSales).Error
	return err
}
