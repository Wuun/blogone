package serializer

import (
	"github.com/jinzhu/gorm"
)

//ArticleModel is the gorm model of article
type ArticleModel struct {
	gorm.Model
	Tag   string `json:"tag"`
	Title string `json:"title"`
}

//ArticleContent is the gorm model of article`s content
type ArticleContent struct {
	ArtModel   ArticleModel `json:"art_model"`
	ArtModelID uint         `json:"art_model_id"`
	Content    string       `gorm:"type:text;" json:"content"`
}
