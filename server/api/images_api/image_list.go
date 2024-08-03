package imagesapi

import (
	"Blog/models"
	"Blog/models/res"
	"Blog/utils"

	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageList(c *gin.Context) {
	imageList := []models.BannerModel{}

	count, list, err := utils.MakePagination(imageList, c)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
	} else {
		res.OKWithList(count, list, c)
	}

}
