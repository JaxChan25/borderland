package model

import (
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
