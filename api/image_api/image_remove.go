package image_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Tags Image
// @Summary 删除图片
// @Description 删除图片
// @Accept  json
// @Produce json
// @Param some_id body models.RemoveList false "要删除的图片id列表"
// @Router /images [delete]
func (ImagesApi) ImageRemove(c *gin.Context) {
	var removeRequest models.RemoveList
	err := c.ShouldBindJSON(&removeRequest)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}
	waitRemove := []models.BannerModel{}
	count := global.DB.Where("id in (?)", removeRequest.IDList).Find(&waitRemove).RowsAffected
	fmt.Println(waitRemove)

	err = global.DB.Delete(&models.BannerModel{}, removeRequest.IDList).Error
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("删除失败%v", err.Error()), c)
	} else {
		res.OKWithMessage(fmt.Sprintf("删除成功,共删除%d条", count), c)
	}
}
