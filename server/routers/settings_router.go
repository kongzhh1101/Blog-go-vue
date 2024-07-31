package routers

import (
	"Blog/api"
)

func (r RouterGroup) SettingsRouter() {
	SettingsApi := api.ApiGroupApp.SettingsApi
	r.GET("settings", SettingsApi.SettingsApiView)
}
