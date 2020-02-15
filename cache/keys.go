package cache

import (
	"fmt"
	"strconv"
)

const (

	//TotalRankKey 总排行
	TotalRankKey = "rank:total"
)

//redis中key的原则
//view:article:1 -> 150
//表示id为1的article的浏览量为150

//ArticleViewKey 生成Article浏览的Key
func ArticleViewKey(id uint) string {
	return fmt.Sprintf("view:article:%s", strconv.Itoa(int(id)))
}

//ArticleLikeKey 生成Article喜欢的Key
func ArticleLikeKey(id uint) string {
	return fmt.Sprintf("like:article:%s", strconv.Itoa(int(id)))
}
