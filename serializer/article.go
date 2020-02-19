package serializer

import "borderland/model"

// Article 文章序列化器
type Article struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	Content      string `json:"content"`
	Catalog      string `json:"catalog"`
	View         uint64 `json:"view"`
	Like         uint64 `json:"like"`
	CreatedAt    string `json:"created_at"`
}

// BuildArticle 序列化文章
func BuildArticle(article model.Article, content string) Article {

	return Article{
		ID:           article.ID,
		Title:        article.Title,
		Introduction: article.Introduction,
		Content:      content,
		Catalog:      article.Catalog,
		View:         article.ViewNumber(),
		Like:         article.LikeNumber(),
		CreatedAt:    article.CreatedAt.Format("2006-01-02 15:04"),
	}

}

// BuildArticleResponse 序列化单篇文章响应
func BuildArticleResponse(article model.Article, content string) Article {
	return BuildArticle(article, content)
}

// BuildArticlesResponse 序列化多篇文章响应
func BuildArticlesResponse(items []model.Article) (articles []Article) {

	for _, item := range items {
		articles = append(articles, BuildArticle(item, ""))
	}
	return articles

}
