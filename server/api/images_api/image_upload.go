package imagesapi

import (
	"Blog/global"
	"Blog/models/res"
	"fmt"
	"io/fs"
	"os"

	"github.com/gin-gonic/gin"
)

type UploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}

func (ImagesApi) UploadImage(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}
	formList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
	}

	basePath := global.Config.Upload.Path
	_, err = os.Stat(basePath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
		}
	}

	var resList []UploadResponse
	for _, file := range formList {
		filepath := fmt.Sprintf("./%s/%s", basePath, file.Filename)
		filesize := float64(file.Size) / float64(1024*1024)
		maxsize := float64(global.Config.Upload.Size)
		if filesize > maxsize {
			resList = append(resList, UploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Message:   fmt.Sprintf("当前图片过大，图片大小必须小于%.0fM，当前图片大小为%.2fM", maxsize, filesize),
			})
			continue
		}
		err = c.SaveUploadedFile(file, filepath)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
		}
		resList = append(resList, UploadResponse{
			FileName:  file.Filename,
			IsSuccess: true,
			Message:   "图片上传成功",
		})
	}

	res.OKWithData(resList, c)
}
