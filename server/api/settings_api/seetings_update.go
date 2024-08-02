package settingsapi

import (
	"Blog/config"
	"Blog/core"
	"Blog/global"
	"Blog/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsUpdate(c *gin.Context) {
	var cr config.SiteInfo
	if err := c.ShouldBindJSON(&cr); err != nil {
		global.Logger.Error(err)
		res.FailWithCode(res.ArguementError, c)
		return
	}

	global.Config.SiteInfo = cr
	err := core.Updateconfig()
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("更新配置失败", c)
		return
	}

	res.OKWithMessage("更新配置成功", c)
}
