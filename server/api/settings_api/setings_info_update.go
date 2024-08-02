package settingsapi

import (
	"Blog/config"
	"Blog/core"
	"Blog/global"
	"Blog/models/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) UpdateSettingsInfo(c *gin.Context) {
	var uri SettingsUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	settingName := uri.Name
	switch settingName {
	case "site":
		var cr config.SiteInfo
		if err := c.ShouldBindJSON(&cr); err != nil {
			global.Logger.Error(err)
			res.FailWithCode(res.ArguementError, c)
			return
		}

		global.Config.SiteInfo = cr
	case "email":
		var cr config.Email
		if err := c.ShouldBindJSON(&cr); err != nil {
			global.Logger.Error(err)
			res.FailWithCode(res.ArguementError, c)
			return
		}

		global.Config.Email = cr
	case "qiniu":
		var cr config.QiNiu
		if err := c.ShouldBindJSON(&cr); err != nil {
			global.Logger.Error(err)
			res.FailWithCode(res.ArguementError, c)
			return
		}

		global.Config.QiNiu = cr
	case "qq":
		var cr config.QQ
		if err := c.ShouldBindJSON(&cr); err != nil {
			global.Logger.Error(err)
			res.FailWithCode(res.ArguementError, c)
			return
		}

		global.Config.QQ = cr
	case "jwt":
		var cr config.JWT
		if err := c.ShouldBindJSON(&cr); err != nil {
			global.Logger.Error(err)
			res.FailWithCode(res.ArguementError, c)
			return
		}

		global.Config.JWT = cr
	}

	err = core.UpdateConfig()
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("更新配置失败", c)
		return
	}

	res.OKWithMessage("更新配置成功", c)
}
