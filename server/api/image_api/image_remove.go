package image_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageRemove(c *gin.Context) {
	var removeRequest models.RemoveList
	err := c.ShouldBindJSON(&removeRequest)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}
	waitRemove := []models.BannerModel{}
	count := global.DB.Where("id in (?)", removeRequest.IDList).Find(&waitRemove).RowsAffected

	global.DB.Delete(&waitRemove)
	res.OKWithMessage(fmt.Sprintf("删除成功,共删除%d条", count), c)
}
