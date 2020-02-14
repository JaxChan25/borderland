package service

import (
	"borderland/cache"
	"borderland/model"
	"borderland/serializer"
	"fmt"
	"strings"
)

// TotalRankService 获得文章总排行top10的服务
type TotalRankService struct {
}

//Get 获取top10
func (service *TotalRankService) Get() serializer.Response {
	var articles []model.Article

	//从redis读取点击前十的视频
	arts, _ := cache.RedisClient.ZRevRange(cache.TotalRankKey, 0, 9).Result()

	order := fmt.Sprintf("FIELD(id,%s)", strings.Join(arts, ","))
	err := model.DB.Where("id in (?)", arts).Order(order).Find(&articles).Error

	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "redis数据库连接错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticlesResponse(articles),
	}
}
