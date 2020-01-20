package services

import (
	"mysql_api/databases"
	"mysql_api/structs/models"
)

func QueryProducts(status int) ([]*models.Product, error) {
	products := []*models.Product{}
	err := databases.Db.Model(&models.Product{}).
		Where(map[string]interface{}{"status": status}).
		Limit(models.EachQueryProductsNum).
		Find(&products).
		Error
	return products, err
}
