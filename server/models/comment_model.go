package models

type CommentModel struct {
	MODEL
	SubComments     *[]CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`
	ParentComment   *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"parent_comment"`
	ParentCommentID *uint           `json:"parent_comment_id"`
	Content         string          `gorm:"size:256" json:"content"`
	LikeCount       int             `gorm:"size:8;default:0" json:"like_count"`
	CommentCount    int             `gorm:"size:8;default:0" json:"comment_count"`
	Article         ArticleModel    `gorm:"foreignKey:ArticleID" json:"-"`
	ArticleID       uint            `json:"article_id"`
	User            UserModel       `json:"user"`
	UserID          uint            `json:"user_id"`
}
