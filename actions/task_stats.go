package actions

import (
	"github.com/gin-gonic/gin"
	"mysql_api/databases"
	"mysql_api/structs/models"
	"net/http"
)

type StatsResponse struct {
	UnSearchCount      int `json:"un_search_count"`
	SearchingCount     int `json:"searching_count"`
	SearchSucceedCount int `json:"search_succeed_count"`
	SearchFailedCount  int `json:"search_failed_count"`
	NoNeedSearchCount  int `json:"no_need_search_count"`
}

// 查询任务状态
func TaskStatsGet(c *gin.Context) {
	var err error
	var unSearchCount int
	var searchingCount int
	var searchSucceedCount int
	var searchFailedCount int
	var noNeedSearchCount int

	err = databases.Db.Model(models.Product{}).Where("status=?", models.StatusUnSearched).Count(&unSearchCount).Error
	err = databases.Db.Model(models.Product{}).Where("status=?", models.StatusSearching).Count(&searchingCount).Error
	err = databases.Db.Model(models.Product{}).Where("status=?", models.StatusSearchSucceed).Count(&searchSucceedCount).Error
	err = databases.Db.Model(models.Product{}).Where("status=?", models.StatusSearchFailed).Count(&searchFailedCount).Error
	err = databases.Db.Model(models.Product{}).Where("status=?", models.StatusNoNeedSearch).Count(&noNeedSearchCount).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": models.CodeQueryDataFailed,
			"msg":  "任务状态查询失败",
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": models.CodeQueryDataSucceed,
		"msg":  "任务状态查询成功",
		"data": &StatsResponse{
			UnSearchCount:      unSearchCount,
			SearchingCount:     searchingCount,
			SearchSucceedCount: searchSucceedCount,
			SearchFailedCount:  searchFailedCount,
			NoNeedSearchCount:  noNeedSearchCount,
		},
	})
}
