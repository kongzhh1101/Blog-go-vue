package api

import (
	imagesapi "Blog/api/images_api"
	settingsapi "Blog/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settingsapi.SettingsApi
	ImagesApi   imagesapi.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
