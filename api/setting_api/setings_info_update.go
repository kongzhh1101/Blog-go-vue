package setting_api

import (
	"Blog/config"
	"Blog/core"
	"Blog/global"
	"Blog/models/res"

	"github.com/gin-gonic/gin"
)

// @Tags 设置信息
// @Summary 跟新设置
// @Description 根据传入的种类，和数据，修改对应设置的信息
// @Accept  json
// @Produce json
// @Param data body config.Config false "设置种类"
// @Param name path string false "设置种类"
// @Router /settings/{name} [put]/
func (SettingsApi) UpdateSettingsInfo(c *gin.Context) {
	var uri SettingsUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithMessage("绑定Uri错误", c)
		return
	}
	settingName := uri.Name
	switch settingName {
	case "site":
		var cr config.SiteInfo
		if err := c.ShouldBindJSON(&cr); err != nil {
			res.FailWithCode(res.ArguementError, c)
			return
		}

		global.Config.SiteInfo = cr
	case "email":
		var cr config.Email
		if err := c.ShouldBindJSON(&cr); err != nil {
			res.FailWithCode(res.ArguementError, c)
			return
		}

		global.Config.Email = cr
	case "qiniu":
		var cr config.QiNiu
		if err := c.ShouldBindJSON(&cr); err != nil {
			res.FailWithCode(res.ArguementError, c)
			return
		}

		global.Config.QiNiu = cr
	case "qq":
		var cr config.QQ
		if err := c.ShouldBindJSON(&cr); err != nil {
			res.FailWithCode(res.ArguementError, c)
			return
		}

		global.Config.QQ = cr
	case "jwt":
		var cr config.JWT
		if err := c.ShouldBindJSON(&cr); err != nil {
			res.FailWithCode(res.ArguementError, c)
			return
		}

		global.Config.JWT = cr
	}

	err = core.UpdateConfig()
	if err != nil {
		res.FailWithMessage("更新配置失败", c)
		return
	}

	res.OKWithMessage("更新配置成功", c)
}
