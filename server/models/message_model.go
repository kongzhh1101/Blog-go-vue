package models

type MessageModel struct {
	MODEL
	SendUserID       uint      `gorm:"primaryKey" json:"send_user_id"`
	SendUser         UserModel `gorm:"foreignKey:SendUserID" json:"send_user"`
	SendUserNickName string    `gorm:"size:42" json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`

	ReceiverUserID       uint      `gorm:"primaryKey" json:"receiver_user_id"`
	ReceiverUser         UserModel `gorm:"foreignKey:ReceiverUserID" json:"receiver_user"`
	ReceiverUserNickName string    `gorm:"size:42" json:"receiver_user_nick_name"`
	ReceiverUserAvatar   string    `json:"recive_user_avatar"`
	IsRead               bool      `gorm:"default:false" json:"is_read"`
	Content              string    `json:"content"`
}
