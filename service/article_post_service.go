package service

import (
	"borderland/model"
	"borderland/serializer"
)

// ArticlePostService 管理文章上传的服务
type ArticlePostService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Catalog string `form:"catalog" json:"catalog"`
}

// valid 验证表单
func (service *ArticlePostService) valid() *serializer.Response {

	count := 0
	model.DB.Model(&model.Article{}).Where("title = ?", service.Title).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "该文章已经上传",
		}
	}
	return nil
}

// Post 用于文章上传
func (service *ArticlePostService) Post() serializer.Response {

	article := model.Article{
		Title:   service.Title,
		Catalog: service.Catalog,
	}

	if err := service.valid(); err != nil {
		return *err
	}

	// 将文章写入数据库
	if err := model.DB.Create(&article).Error; err != nil {
		return serializer.ParamErr("写入文章失败", err)
	}

	return serializer.BuildArticleResponse(article, "")
}
