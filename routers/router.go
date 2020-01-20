package routers

import (
	"github.com/gin-gonic/gin"
	"mysql_api/actions"
)

var r *gin.Engine

func Load() *gin.Engine {
	r.POST("insert/product-info", actions.InsertProductInfoPost)
	r.POST("insert/price-sales", actions.InsertPriceSalesPost)
	r.GET("query/products", actions.ProductsGet)
	r.POST("update/product-status", actions.UpdateProductStatusPost)
	r.GET("util/task-stats", actions.TaskStatsGet)

	return r
}

func init() {
	r = gin.Default()
}
