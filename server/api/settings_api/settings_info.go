package settingsapi

import (
	"Blog/global"
	"Blog/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsApiView(c *gin.Context) {
	res.OKWithData(global.Config.SiteInfo, c)
}
