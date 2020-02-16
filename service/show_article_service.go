package service

import (
	"borderland/model"
	"borderland/serializer"
	"borderland/util"
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
			Msg:   "文章查询失败",
			Error: err.Error(),
		}
	}

	//文章点击量加一
	article.Addview()

	title := article.Title
	objectName := "article/" + title + ".md"
	content, err := util.GetOssStream(objectName)

	return serializer.Response{
		Data: serializer.BuildArticleResponse(article, content),
	}
}
