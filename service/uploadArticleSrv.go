package service

import (
	"bblog/conf"
	"bblog/serializer"
	"github.com/pborman/uuid"
)

type UploadArticleSrv struct {
	Title   string `form:"title"`
	Tag     string `form:"tag"`
	Content string `form:"article_content"`
}

func (service *UploadArticleSrv) Upload() *serializer.Response {
	ud := uuid.New()
	articleModel := &serializer.ArticleModel{
		UUID:  ud,
		Tag:   service.Tag,
		Title: service.Title,
	}
	articleContentModel := &serializer.ArticleContent{
		ArticleModelID: ud,
		Content:        service.Content,
	}

	if err := conf.MYSQL_CONNECT.Create(articleModel).Error; err != nil {
		return &serializer.Response{
			Status: 40002,
			Msg:    "can't not insert this article into server.",
			Error:  err,
		}
	}

	if err := conf.MYSQL_CONNECT.Create(articleContentModel).Error; err != nil {
		conf.MYSQL_CONNECT.Delete(articleModel).Where("uuid = ?", ud)
		return &serializer.Response{
			Status: 40002,
			Msg:    "can't not insert this article into server.",
			Error:  err,
		}
	}

	return &serializer.Response{
		Status: 0,
		Msg:    "success",
	}
}
