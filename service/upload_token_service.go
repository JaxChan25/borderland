package service

import (
	"borderland/serializer"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

// UploadTokenService 获得阿里云上传oss token的服务
type UploadTokenService struct {
	Filename string `form:"filename" json:"filename"`
}

// Post 获得阿里云上传oss token
func (service *UploadTokenService) Post() serializer.Response {
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACESS_KEY"), os.Getenv("OSS_KEY_SECRET"))
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeEnv,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}
	//获得存储空间
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeEnv,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	//上传可选参数 - 签名直传
	options := []oss.Option{
		oss.ContentType("image/jpeg"),
	}

	key := "avatar/" + uuid.Must(uuid.NewRandom()).String() + ".jpg"

	//进行签名直传
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeEnv,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeEnv,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"filename": key,
			"putapi":   signedPutURL,
			"getapi":   signedGetURL,
		},
	}

}
