package models

import (
	"Blog/global"
	"Blog/models/ctype"
	"os"

	"gorm.io/gorm"
)

type BannerModel struct {
	MODEL
	Path     string         `json:"path"`
	Hash     string         `json:"hash"`
	Name     string         `gorm:"size:38" json:"name"`
	Location ctype.Location `gorm:"default:1" json:"location"`
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	// 判断文件存储的位置，本地or云端
	if b.Location == ctype.Local {

		err = os.Remove(b.Path)
		if err != nil {
			global.Logger.Error(err)
			return err
		}

	}
	return nil
}
