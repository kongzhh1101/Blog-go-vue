package image_api

import (
	"Blog/models"
	"Blog/models/res"
	service "Blog/service/common"

	"github.com/gin-gonic/gin"
)

// @Tags Image
// @Summary 查询图片
// @Description 获取图片列表
// @Produce json
// @Param page query integer false "页码"
// @Param limit query integer false "每页显示所少"
// @Router /images [get]
func (ImagesApi) ViewImage(c *gin.Context) {

	// 获取要返回的结果切片
	imageSlice := []models.BannerModel{}
	count, resultSlice, err := service.MakePagination(imageSlice, c)

	if err != nil {
		res.FailWithMessage(err.Error(), c)
	} else {
		//获取成功，将信息返回前端
		res.OKWithSlice(count, resultSlice, c)
	}

}
