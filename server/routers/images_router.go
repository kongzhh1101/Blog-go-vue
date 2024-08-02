package routers

import "Blog/api"

func (r RouterGroup) ImagesRouter() {
	ImagesApi := api.ApiGroupApp.ImagesApi
	r.POST("images", ImagesApi.UploadImage)
}
