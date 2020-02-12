package serializer

import "borderland/model"

// Article 文章序列化器
type Article struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Catalog   string `json:"catalog"`
	CreatedAt int64  `json:"created_at"`
}

// BuildArticle 序列化文章
func BuildArticle(article model.Article, content string) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   content,
		Catalog:   article.Catalog,
		CreatedAt: article.CreatedAt.Unix(),
	}
}

// BuildArticleResponse 序列化文章响应
func BuildArticleResponse(article model.Article, content string) Response {
	return Response{
		Data: BuildArticle(article, content),
	}
}
