package serializer

import (
	"time"
)

//Article is the gorm model of article
type Article struct {
	CreateTime time.Time `json:"create_time"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
}

type ArticleList struct {
	Title string `json:"title"`
	Tag   []Tag  `json:"tag"`
}

type Tag struct {
	ArticleTag string `json:"article_tag"`
}
