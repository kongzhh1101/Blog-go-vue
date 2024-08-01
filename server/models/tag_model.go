package models

type TagModel struct {
	MODEL
	Name     string         `gorm:"size:16" json:"name"`
	Articles []ArticleModel `gorm:"many2many:article_tag" json:"-"`
}
