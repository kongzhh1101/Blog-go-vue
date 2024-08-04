package api

import (
	"Blog/api/image_api"
	"Blog/api/setting_api"
)

type ApiGroup struct {
	SettingsApi setting_api.SettingsApi
	ImagesApi   image_api.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
