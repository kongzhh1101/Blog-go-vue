package models

import (
	"Blog/global"
	"Blog/models/ctype"
	"fmt"
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
	fmt.Println("xxxxxxxxx")
	fmt.Println(b.Path)
	if b.Location == ctype.Local {
		fmt.Println("aaaaaa")
		err = os.Remove(b.Path)
		if err != nil {
			global.Logger.Error(err)
			return err
		}
	}
	return nil
}
