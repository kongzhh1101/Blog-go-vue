package routers

import (
	"Blog/api"
)

func (r RouterGroup) SettingRouter() {
	SettingsApi := api.ApiGroupApp.SettingsApi
	r.GET("settings/:name", SettingsApi.GetSettingsInfo)
	r.PUT("settings/:name", SettingsApi.UpdateSettingsInfo)
}
