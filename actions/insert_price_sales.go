package actions

import (
	"github.com/gin-gonic/gin"
	"mysql_api/services"
	"mysql_api/structs"
	"mysql_api/structs/models"
	"net/http"
	"strings"
)

// 插入商品每天的价格和销量
func InsertPriceSalesPost(c *gin.Context) {
	priceSales := &structs.ProductPriceSales{}
	if err := c.BindJSON(priceSales); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": models.CodeRequestFormError,
			"msg":  "请求格式不正确",
			"err":  err.Error(),
		})
		return
	}

	if err := services.InsertPriceSales(priceSales); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusOK, gin.H{
				"code": models.CodeInsertDataRepeated,
				"msg":  "重复插入数据",
				"err":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"code": models.CodeInsertDataFailed,
			"msg":  "插入数据失败",
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": models.CodeInsertDataSucceed,
		"msg":  "插入数据成功",
	})
}
