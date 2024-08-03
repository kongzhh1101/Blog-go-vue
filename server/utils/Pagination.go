package utils

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"

	"github.com/gin-gonic/gin"
)

func MakePagination[T any](model []T, c *gin.Context) (count any, list []T, err error) {
	count = global.DB.Find(&model).RowsAffected

	var pageInfo models.Page
	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		res.FailWithCode(res.ArguementError, c)
		return
	}

	if pageInfo.Limit == 0 {
		err = global.DB.Find(&model).Error
	} else {
		page := pageInfo.Page
		offset := pageInfo.Limit * (page - 1)
		if offset < 0 {
			offset = 0
		}
		err = global.DB.Limit(pageInfo.Limit).Offset(offset).Find(&model).Error
	}

	return count, model, err
}
