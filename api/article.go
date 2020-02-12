package api

import (
	"borderland/service"

	"github.com/gin-gonic/gin"
)

// ArticlePost 提交markdown格式的article
func ArticlePost(c *gin.Context) {

	service := service.ArticlePostService{}

	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
