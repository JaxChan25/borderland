package util

import (
	"io/ioutil"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//GetOssStream 流式下载
func GetOssStream(objectName string) (string, error) {
	// 创建OSSClient实例。
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACESS_KEY"), os.Getenv("OSS_KEY_SECRET"))
	if err != nil {
		return "", err
	}

	// 获取存储空间。
	bucket, err := client.Bucket("borderland")
	if err != nil {
		return "", err
	}

	// 下载文件到流。
	body, err := bucket.GetObject(objectName)
	if err != nil {
		return "", err
	}

	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
