package global

import (
	"Blog/config"

	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
