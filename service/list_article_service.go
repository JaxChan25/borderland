package service

import (
	"borderland/model"
	"borderland/serializer"
)

// ListArticleService 文章列表的服务
type ListArticleService struct {
	Limit   int    `form:"limit"`
	Start   int    `form:"start"`
	Catalog string `form:"catalog"`
}

// List 展示文章列表
func (service *ListArticleService) List() serializer.Response {

	var articles []model.Article
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if service.Catalog == "" {
		//文章总数
		if err := model.DB.Model(model.Article{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库查询错误",
				Error: err.Error(),
			}
		}

		if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&articles).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库查询错误",
				Error: err.Error(),
			}
		}

	} else { //catalog不为空

		if err := model.DB.Model(model.Article{}).Where("catalog = ?", service.Catalog).Count(&total).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库查询错误",
				Error: err.Error(),
			}
		}

		if err := model.DB.Limit(service.Limit).Offset(service.Start).Where("catalog = ?", service.Catalog).Find(&articles).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库查询错误",
				Error: err.Error(),
			}
		}

	}

	return serializer.BuildListResponse(serializer.BuildArticlesResponse(articles), uint(total))

}
