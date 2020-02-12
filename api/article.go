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

// ShowArticle 取得单个article文件
func ShowArticle(c *gin.Context) {
	service := service.ShowArticleService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListArticle 取得所有article文件
func ListArticle(c *gin.Context) {
	service := service.ListArticleService{}
	res := service.List()
	c.JSON(200, res)
}
