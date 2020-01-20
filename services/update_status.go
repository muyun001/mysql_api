package services

import (
	"mysql_api/databases"
	"mysql_api/structs"
	"mysql_api/structs/models"
)

// 更新状态码
func UpdateStatus(codeStatus *[]structs.UpdateStatusRequest) error {
	for _, singleData := range *codeStatus {
		err := databases.Db.Model(&models.Product{}).
			Where("product_code=?", singleData.ProductCode).
			Update("status", singleData.Status).
			Error
		if err != nil {
			return err
		}
	}
	return nil
}
