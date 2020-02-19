package service

import (
	"borderland/model"
	"borderland/serializer"
)

// ListCatalogService 文章列别的服务
type ListCatalogService struct {
}

// List 展示文章类别
func (service *ListCatalogService) List() serializer.Response {

	catalogs := make([]string, 0, 4)

	rows, err := model.DB.Raw("SELECT DISTINCT catalog FROM articles").Rows()
	defer rows.Close()

	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}

	for rows.Next() {
		var tmp string
		rows.Scan(&tmp)
		catalogs = append(catalogs, tmp)
	}

	return serializer.BuildListResponse(catalogs, uint(len(catalogs)))
}
