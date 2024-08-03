package models

import "time"

type MODEL struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"`
	UpdateAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

type Page struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
