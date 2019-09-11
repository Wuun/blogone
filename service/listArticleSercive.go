package service

import (
	"bblog/conf"
	"bblog/serializer"
)

//HomePageServ is service use in home page.
type HomePageServ struct {
	ArticleSet []serializer.ArticleModel
	pageCount  int
	Page       []int
	Limit      int
	offset     int
}

//List is use to list article.
func (hpSrv *HomePageServ) List() {
	var (
		articelMd []serializer.ArticleModel
	)
	conf.MYSQL_CONNECT.Order("created_at desc").Find(&articelMd)
	if hpSrv.Limit == 0 {
		hpSrv.Limit = 4
	}
	hpSrv.setArticleSet(articelMd)
	hpSrv.pageCount = setPageCount(len(articelMd), hpSrv.Limit)
	hpSrv.setPage()
}

//SetLimit set HomePageServ's limit.
func (hpSrv *HomePageServ) SetLimit(limit int) {
	hpSrv.Limit = limit
}

//SetOffset set the offset of hpsrv.
func (hpSrv *HomePageServ) SetOffset(offset int) {
	hpSrv.offset = offset
}

func setPageCount(a int, b int) int {
	tem := a / b
	tem2 := float32(a) / float32(b)
	if float32(tem) < tem2 {
		return tem + 1
	}
	return tem
}

func (hpSrv *HomePageServ) setPage() {
	for i := 1; i < hpSrv.pageCount+1; i++ {
		hpSrv.Page = append(hpSrv.Page, i)
	}
}

func (hpSrv *HomePageServ) setArticleSet(articelMd []serializer.ArticleModel) {
	if len(articelMd) < hpSrv.Limit {
		hpSrv.ArticleSet = articelMd
	} else {
		hpSrv.ArticleSet = articelMd[hpSrv.offset : hpSrv.Limit+hpSrv.offset]
	}
}
