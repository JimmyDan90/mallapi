package api

import (
	"github.com/gin-gonic/gin"
	"mallapi/response"
	"mallapi/service"
)

var upload service.WebUploadService
// WebFileUpload 图片文件上传
func WebFileUpload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.Failed("图片上传失败", c)
		return
	}
	fileSize := fileHeader.Size
	url := upload.UploadToQiniu(file, fileSize)
	response.Success("图片上传成功", url, c)
}