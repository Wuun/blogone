package service

import (
	"blogone/conf"
	"blogone/serializer"
)

type ArticleCommentSrv struct {
	ArticleID string		`form:"article_id";json:"article_id"`
	Author string			`form:"author";json:"author"`
	Comment string			`form:"comment";json:"comment"`
}

func (comment *ArticleCommentSrv)Add()*serializer.Response{
	commentModel := &serializer.Comment{
		ArticleID:comment.ArticleID,
		Text:comment.Comment,
		Author:comment.Author,
	}
	if err := conf.MYSQL_CONNECT.Create(commentModel).Error;err != nil {
		return &serializer.Response{
			Status:40001,
			Error:err,
		}
	}
	return &serializer.Response{
		Status:0,
		Msg:"comment successfully.",
	}
}