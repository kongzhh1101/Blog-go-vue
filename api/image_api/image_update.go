package image_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"

	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

// @Tags Image
// @Summary 更新图片
// @Description 更改图片名字
// @Accept  json
// @Produce json
// @Param id body ImageUpdateRequest false "要修改的图片id"
// @Router /images [put]
func (ImagesApi) UpdateImage(c *gin.Context) {

	var cr ImageUpdateRequest

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var imageModel models.BannerModel

	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("文件不存在", c)
		return
	}

	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OKWithMessage("图片名修改成功", c)
}
