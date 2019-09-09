package serializer

//ArticleModel is the gorm model of article
type ArticleModel struct {
	UUID string `gorm:"PRIMARY_KEY;UNIQUE;type:varchar(30)"`
	Tag   string `json:"tag"`
	Title string `json:"title"`
}

//ArticleContent is the gorm model of article`s content
type ArticleContent struct {
	ArticleModel   ArticleModel `json:"art_model";gorm:"foreignkey:ArticleModelID"`
	ArticleModelID  string 			`gorm:"PRIMARY_KEY"`
	Content    string       `gorm:"type:text;" json:"content"`
}

//Comment is the comment of a specific article.
type Comment struct {
	ArticleContent ArticleContent		`gorm:"foreignkey:ArticleID"`
	ArticleID string		`gorm:"PRIMARY_KEY"`
	Text string		`gorm:"VARCHAR(60)"`
	Author string	`gorm:"VARCHAR(20)"`
}