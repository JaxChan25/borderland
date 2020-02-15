package model

import (
	"borderland/cache"
	"strconv"

	"github.com/jinzhu/gorm"
)

// Article 文章模型
type Article struct {
	gorm.Model
	Title   string
	Catalog string
}

// GetArticle 用ID获取文章
func GetArticle(ID interface{}) (Article, error) {
	var article Article
	result := DB.First(&article, ID)
	return article, result.Error
}

//ViewNumber 文章游览累计数
func (article *Article) ViewNumber() uint64 {

	countStr, _ := cache.RedisClient.Get(cache.ArticleViewKey(article.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//Addview 文章游览+1
func (article *Article) Addview() {

	//增加视频点击数
	cache.RedisClient.Incr(cache.ArticleViewKey(article.ID))

	//增加排行榜里面的视频点击数
	cache.RedisClient.ZIncrBy(cache.TotalRankKey, 1, strconv.Itoa(int(article.ID)))

}

//LikeNumber 文章喜欢的累计数
func (article *Article) LikeNumber() uint64 {

	countStr, _ := cache.RedisClient.Get(cache.ArticleLikeKey(article.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//AddLike 文章喜欢+1
func (article *Article) AddLike() {

	//增加视频点击数
	cache.RedisClient.Incr(cache.ArticleLikeKey(article.ID))

}
