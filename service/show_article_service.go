package service

import (
	"borderland/model"
	"borderland/serializer"
	"io/ioutil"
	"os"
)

// ShowArticleService 文章详情的服务
type ShowArticleService struct {
}

// Show 展示视频
func (service *ShowArticleService) Show(id string) serializer.Response {

	var article model.Article
	err := model.DB.First(&article, id).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频查询失败",
			Error: err.Error(),
		}
	}

	title := article.Title
	filePath := os.Getenv("ARTICLE_ROOT") + title + ".md"

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "markdown文件打开失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticleResponse(article, string(content)),
	}
}
