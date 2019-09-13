package service

import (
	"bblog/conf"
	"bblog/serializer"
)

type ModifyArticleSrv struct {
	Title       string
	Description string
	Content     string
}

func (modSrv *ModifyArticleSrv) FindInformation(articleID string) error {
	var (
		articleModel   serializer.ArticleModel
		articleContent serializer.ArticleContent
	)
	if err := conf.MYSQL_CONNECT.Where("uuid = ?", articleID).Find(&articleModel).Error; err != nil {
		return err
	}
	if err := conf.MYSQL_CONNECT.Where("article_model_id = ?", articleID).Find(&articleContent).Error; err != nil {
		return err
	}

	modSrv.Title = articleModel.Title
	modSrv.Description = articleModel.Description
	modSrv.Content = articleContent.Content
	return nil
}

func (modSrv *ModifyArticleSrv) Modify(articleID string) error {
	var (
		articleModel   serializer.ArticleModel
		articleContent serializer.ArticleContent
	)
	articleModel.UUID = articleID
	articleContent.ArticleModelID = articleID
	if err := conf.MYSQL_CONNECT.Model(&articleModel).Updates(map[string]interface{}{
		"title":       modSrv.Title,
		"description": modSrv.Description,
	}).Error; err != nil {
		return nil
	}
	if err := conf.MYSQL_CONNECT.Model(&articleContent).Updates(map[string]interface{}{
		"content": modSrv.Content,
	}).Error; err != nil {
		return nil
	}
	return nil
}
