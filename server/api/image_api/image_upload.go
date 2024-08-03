package image_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"Blog/utils"
	"fmt"
	"io/fs"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

var ImageWhiteList = []string{
	".jpg",
	".jpeg",
	".png",
	".gif",
	".bmp",
	".svg",
	".webp",
	".ico",
	".tif",
	".tiff",
}

// UploadResponse 图片上传响应结构体
// 用于返回图片上传的结果信息
type UploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Message   string `json:"message"`    // 消息描述
}

// UploadImage 上传图片接口处理函数
// 该函数处理图片上传请求，验证图片文件并保存到指定路径
// 参数:
//
//	c *gin.Context - Gin的上下文对象，用于处理HTTP请求和响应
func (ImagesApi) UploadImage(c *gin.Context) {
	// 解析表单数据
	form, err := c.MultipartForm()
	if err != nil {
		// 解析表单数据失败，返回错误信息
		res.FailWithMessage(err.Error(), c)
	}
	// 获取文件列表
	FileList, ok := form.File["images"]
	if !ok {
		// 文件不存在，返回错误信息
		res.FailWithMessage("不存在的文件", c)
	}

	// 获取上传路径
	basePath := global.Config.Upload.Path
	// 检查路径是否存在
	_, err = os.Stat(basePath)
	if os.IsNotExist(err) {
		// 路径不存在，创建路径
		err := os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			// 创建路径失败，返回错误信息
			res.FailWithMessage(err.Error(), c)
		}
	}

	// 初始化结果列表
	var resList []UploadResponse
	// 遍历文件列表
	for _, file := range FileList {
		// 将文件名转换为小写
		fileSuffix := strings.ToLower(path.Ext(file.Filename))
		// 如果文件后缀不在白名单中
		if !slices.Contains(ImageWhiteList, fileSuffix) {
			// 如果文件格式不被支持，则将文件名、是否成功、错误信息添加到resList中
			resList = append(resList, UploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Message:   "不支持的文件格式",
			})
			// 继续处理下一个文件
			continue
		}

		// 计算上传文件大小
		filesize := float64(file.Size) / float64(1024*1024)

		// 获取最大文件大小限制
		maxsize := float64(global.Config.Upload.Size)

		// 检查文件大小是否超过限制
		if filesize > maxsize {
			// 文件过大，添加到结果列表
			resList = append(resList, UploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Message:   fmt.Sprintf("当前图片过大，图片大小必须小于%.0fM，当前图片大小为%.2fM", maxsize, filesize),
			})
			continue
		}

		// 构造文件路径
		filepath := fmt.Sprintf("./%s/%s", basePath, file.Filename)

		// 上传到数据库
		f, err := file.Open()
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		defer f.Close()

		buffer := make([]byte, file.Size)

		// 将哈希值转换为十六进制字符串
		hashString := utils.MD5(buffer)

		err = global.DB.Where("hash=?", hashString).Take(&models.BannerModel{}).Error
		if err != nil {
			// 没找到
			// 保存文件
			err = c.SaveUploadedFile(file, filepath)
			if err != nil {
				// 保存文件失败，返回错误信息
				res.FailWithMessage(err.Error(), c)
			}

			imageinfo := models.BannerModel{
				Path: filepath,
				Hash: hashString,
				Name: file.Filename,
			}

			global.DB.Create(&imageinfo)

			// 文件上传成功，添加到结果列表
			resList = append(resList, UploadResponse{
				FileName:  file.Filename,
				IsSuccess: true,
				Message:   "图片上传成功",
			})

		} else {
			// 找到了
			resList = append(resList, UploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Message:   "图片已存在",
			})
			continue
		}

	}

	// 返回上传结果
	res.OKWithData(resList, c)
}
