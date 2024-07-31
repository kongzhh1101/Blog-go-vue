package settingsapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsApiView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "test",
	})
}
