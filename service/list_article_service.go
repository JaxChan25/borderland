package service

import (
	"borderland/model"
	"borderland/serializer"
)

// ListArticleService 文章列表的服务
type ListArticleService struct {
}

// List 展示文章列表
func (service *ListArticleService) List() serializer.Response {

	var articles []model.Article

	err := model.DB.Find(&articles).Error

	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticlesResponse(articles),
	}
}
