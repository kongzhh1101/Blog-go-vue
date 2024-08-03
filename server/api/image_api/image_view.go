package image_api

import (
	"Blog/models"
	"Blog/models/res"
	"Blog/utils"

	"github.com/gin-gonic/gin"
)

type GetResponse struct {
	ID       uint   `json:"id"`
	FileName string `json:"file_name"`
	Hash     string `json:"hash"`
}

func (ImagesApi) ViewImage(c *gin.Context) {
	var getList []GetResponse

	imageList := []models.BannerModel{}

	count, list, err := utils.MakePagination(imageList, c)

	for _, v := range list {
		getList = append(getList, GetResponse{
			ID:       v.ID,
			FileName: v.Name,
			Hash:     v.Hash,
		})
	}

	if err != nil {
		res.FailWithMessage(err.Error(), c)
	} else {
		res.OKWithList(count, getList, c)
	}

}
