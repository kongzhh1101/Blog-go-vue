package flag

import (
	"Blog/global"
	"Blog/models"
)

func MakeMigration() {
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectionModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	err = global.DB.AutoMigrate(
		&models.UserModel{},
		&models.ArticleModel{},
		&models.CommentModel{},
		&models.TagModel{},
		&models.BannerModel{},
		&models.MenuModel{},
		&models.MessageModel{},
		&models.AdvertisementModel{},
		&models.FeedbackModel{},
		&models.LoginDataModel{},
	)

	if err != nil {
		global.Logger.Error("[error] 生成数据表失败")
		return
	} else {
		global.Logger.Info("[success] 生成数据表成功")
	}
}
