package service

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mallapi/global"
	"mime/multipart"
)

type WebUploadService struct {
}

// UploadToQiniu 图片上传到七牛云然后返回图片URL
func (u *WebUploadService) UploadToQiniu(file multipart.File, fileSize int64)  string {
	qiniuStorage := global.Config.QiniuStorage
	putPlolicy := storage.PutPolicy{
		Scope: qiniuStorage.Bucket,
	}
	mac := qbox.NewMac(qiniuStorage.AccessKey, qiniuStorage.SecretKey)
	upToken := putPlolicy.UploadToken(mac)
	fmt.Println(upToken)
	cfg := storage.Config{
		Zone: &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS: false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(),&ret,upToken,file,fileSize,&putExtra)
	if err != nil {
		return err.Error()
	}
	url := qiniuStorage.QiniuServer + ret.Key
	return url
}
