package setting_api

import (
	"Blog/global"
	"Blog/models/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// @Tags 设置信息
// @Summary 查看设置信息
// @Description 根据传入的种类，返回对应的设置信息
// @Produce json
// @Param name path string false "设置种类"
// @Router /settings/{name} [get]/
func (SettingsApi) GetSettingsInfo(c *gin.Context) {
	var uri SettingsUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithMessage("绑定Uri错误", c)
		return
	}
	settingName := uri.Name
	switch settingName {
	case "site":
		res.OKWithData(global.Config.SiteInfo, c)
	case "email":
		res.OKWithData(global.Config.Email, c)
	case "jwt":
		res.OKWithData(global.Config.JWT, c)
	case "qq":
		res.OKWithData(global.Config.QQ, c)
	case "qiniu":
		res.OKWithData(global.Config.QiNiu, c)
	default:
		c.JSON(http.StatusNotFound, gin.H{"error": "Setting type not found"})
	}
}
