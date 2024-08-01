package models

import "Blog/models/ctype"

type ArticleModel struct {
	MODEL
	Title         string         `gorm:"size:32" json:"title"`
	Abstract      string         `json:"abstract"`
	Content       string         `json:"content"`
	LookCount     int            `json:"look_count"`
	CommentCount  int            `json:"comment_count"`
	LikeCount     int            `json:"like_count"`
	CollectCount  int            `json:"collect_count"`
	TagModels     []TagModel     `gorm:"many2many:atricle_tag" json:"tag_models"`
	CommentModels []CommentModel `gorm:"foreignKey:Article" json:"-"`
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`
	UserID        uint           `json:"user_id"`
	Category      string         `gorm:"size:28" json:"category"`
	Source        string         `json:"source"`
	Link          string         `json:"link"`
	Banner        BannerModel    `gorm:"foreignKey:BannerID" json:"-"`
	BannerID      uint           `json:"banner_id"`
	NickName      string         `gorm:"size:42" json:"nick_name"`
	BannerPath    string         `json:"banner_path"`
	Tags          ctype.Array    `gorm:"type:string;size:64" json:"tag"`
}
