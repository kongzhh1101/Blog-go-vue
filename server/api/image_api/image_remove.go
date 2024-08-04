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
	fmt.Println(waitRemove)

	err = global.DB.Delete(&waitRemove).Error
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("删除失败%v", err.Error()), c)
	} else {
		res.OKWithMessage(fmt.Sprintf("删除成功,共删除%d条", count), c)
	}
}
