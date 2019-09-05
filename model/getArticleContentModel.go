package model

import "blogone/service"

func GetarticleContentModel(articleId int) (result []byte, err error) {
	articalContent, err := service.GetArticleContentSrv(articleId)
	if err != nil {
		return []byte(""), err
	}

	return []byte(articalContent.Content), nil
}
