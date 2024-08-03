package routers

import "Blog/api"

func (r RouterGroup) ImageRouter() {
	ImagesApi := api.ApiGroupApp.ImagesApi
	r.GET("images", ImagesApi.ViewImage)
	r.POST("images", ImagesApi.UploadImage)
	r.DELETE("images", ImagesApi.ImageRemove)
}
