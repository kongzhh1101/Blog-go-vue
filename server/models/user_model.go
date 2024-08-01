package models

import "Blog/models/ctype"

// UserModel 用户表
type UserModel struct {
	MODEL
	NickName         string             `gorm:"size:36" json:"nick_name"`
	UserName         string             `gorm:"size:36" json:"user_name"`
	Password         string             `gorm:"size:128" json:"password"`
	Avatar           string             `gorm:"size:191" json:"avatar_id"`
	Email            string             `gorm:"size:128" json:"email"`
	Tel              string             `gorm:"size:18" json:"tel"`
	Addr             string             `gorm:"size:64" json:"addr"`
	Token            string             `gorm:"size:64" json:"token"`
	IP               string             `gorm:"size:20" json:"ip"`
	Role             ctype.Role         `gorm:"size:4;default:1" json:"role"`
	SignupSource     ctype.SignupSource `gorm:"type:smallint(6)" json:"singup_source"`
	ArticleModels    []ArticleModel     `gorm:"foreignKey:UserID" json:"-"`
	CollectionModels []ArticleModel     `gorm:"many2many:user_collection;ForeignKey:UserID;joinForeignKey:ArticleID" json:"-"`
}
