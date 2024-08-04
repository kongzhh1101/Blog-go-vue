package models

type FeedbackModel struct {
	MODEL
	Email        string `gorm:"size:64" json:"email"`
	Content      string `gorm:"size:128" json:"content"`
	ReplyContent string `gorm:"size:128" json:"reply_content"`
	IsReply      bool   `json:"is_reply"`
}
