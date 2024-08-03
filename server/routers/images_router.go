package routers

import "Blog/api"

func (r RouterGroup) ImagesRouter() {
	ImagesApi := api.ApiGroupApp.ImagesApi
	r.GET("images", ImagesApi.ImageList)
	r.POST("images", ImagesApi.UploadImage)
}
