package serializer

import "borderland/model"

// Article 文章序列化器
type Article struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Catalog   string `json:"catalog"`
	View      uint64 `json:"view"`
	Like      uint64 `json:"like"`
	CreatedAt string `json:"created_at"`
}

// BuildArticle 序列化文章
func BuildArticle(article model.Article, content string) Article {

	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   content,
		Catalog:   article.Catalog,
		View:      article.ViewNumber(),
		Like:      article.LikeNumber(),
		CreatedAt: article.CreatedAt.Format("2006年1月2日 15:04:05"),
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
