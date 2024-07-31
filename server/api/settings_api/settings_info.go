package settingsapi

import (
	"Blog/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsApiView(c *gin.Context) {
	res.FailWithCode(res.SettingError, c)
}
