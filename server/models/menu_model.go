package models

import "Blog/models/ctype"

type MenuModel struct {
	MODEL
	MenuTitle    string        `gorm:"size:32" json:"menu_title"`
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"`
	Slogen       string        `gorm:"size:64" json:"slogen"`
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`
	AbstractTime int           `json:"abstract_time"`
	Banners      []BannerModel `gorm:"many2many:menu_banner; joinForeignKey:MenuID, joinReferences:BannerID" json:"banners"`
	BannerTime   int           `json:"banner_time"`
	Sort         int           `gorm:"size:10" json:"sort"`
}
