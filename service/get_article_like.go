package service

import (
	"borderland/model"
	"borderland/serializer"
)

//AddArticleLikeService 增加文章点赞数量的服务
type AddArticleLikeService struct {
}

//Add 增加文章点赞数
func (service *AddArticleLikeService) Add(id string) serializer.Response {

	var article model.Article
	err := model.DB.First(&article, id).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "文章查询失败",
			Error: err.Error(),
		}
	}

	//文章点赞数加一
	article.AddLike()

	return serializer.Response{
		Data: serializer.BuildArticleResponse(article, ""),
	}

}
