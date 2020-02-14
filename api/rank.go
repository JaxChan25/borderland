package api

import (
	"borderland/service"

	"github.com/gin-gonic/gin"
)

//TotalRank 获取文章总点击量排行
func TotalRank(c *gin.Context) {

	service := service.TotalRankService{}
	res := service.Get()
	c.JSON(200, res)

}
