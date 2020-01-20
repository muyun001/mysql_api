package actions

import (
	"github.com/gin-gonic/gin"
	"mysql_api/services"
	"mysql_api/structs"
	"mysql_api/structs/models"
	"net/http"
)

// 更新状态码
func UpdateProductStatusPost(c *gin.Context) {
	request := &[]structs.UpdateStatusRequest{}
	if err := c.BindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": models.CodeRequestFormError,
			"msg":  "请求格式不正确",
			"err":  err.Error(),
		})
		return
	}

	if err := services.UpdateStatus(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": models.CodeUpdateDataFailed,
			"msg":  "更新数据失败",
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": models.CodeUpdateDataSucceed,
		"msg":  "更新数据成功",
	})
}
