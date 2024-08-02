package settingsapi

import (
	"Blog/global"
	"Blog/models/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) GetSettingsInfo(c *gin.Context) {
	var uri SettingsUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
