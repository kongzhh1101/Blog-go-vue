package service

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"

	"github.com/gin-gonic/gin"
)

// 用于分页查询，根据前端传回的分页信息，返回对应的结果切片
// 使用泛型，根据传入的模型自动判断
func MakePagination[T any](modelSlice []T, c *gin.Context) (count any, resultSlice []T, err error) {
	// 获取数据库中该模型的数据总数
	count = global.DB.Find(&modelSlice).RowsAffected

	// 获取前端请求内容
	var pageInfo models.Page
	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		res.FailWithCode(res.ArguementError, c)
		return
	}

	// 如果前端请求内容中没有排序，默认根据创建时间降序
	sort := pageInfo.Sort
	if sort == "" {
		sort = "create_at desc"
	}

	// 如果前端请求内容中没有指明每页的限制，默认不分页
	if pageInfo.Limit == 0 {
		err = global.DB.Order(sort).Find(&modelSlice).Error
	} else {
		//如果前端请求内容中没有指明页码或者页码为负数，默认第一页
		page := pageInfo.Page
		offset := pageInfo.Limit * (page - 1)

		if offset < 0 {
			offset = 0
		}

		err = global.DB.Order(sort).Limit(pageInfo.Limit).Offset(offset).Find(&modelSlice).Error
	}

	return count, modelSlice, err
}
