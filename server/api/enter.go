package api

import settingsapi "Blog/api/settings_api"

type ApiGroup struct {
	SettingsApi settingsapi.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
