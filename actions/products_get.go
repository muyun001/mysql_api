package actions

import (
	"github.com/gin-gonic/gin"
	"mysql_api/services"
	"mysql_api/structs/models"
	"net/http"
	"strconv"
)

// 查询未抓取的商品
func ProductsGet(c *gin.Context) {
	statusStr := c.Query("status")

	var err error
	status := models.StatusUnSearched
	if statusStr != "" {
		status, err = strconv.Atoi(statusStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": models.CodeRequestFormError,
				"msg":  "请求格式错误",
				"err":  err.Error(),
			})
			return
		}
	}

	products, err := services.QueryProducts(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": models.CodeQueryDataFailed,
			"msg":  "查询数据失败",
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": models.CodeQueryDataSucceed,
		"msg":  "查询数据成功",
		"data": products,
	})
}
