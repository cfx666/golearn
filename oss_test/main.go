package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func main() {
	// yourBucketName填写存储空间名称。
	bucketName := "mycrowd-01"
	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	objectName := "test/test1.jpg"
	// yourLocalFileName填写本地文件的完整路径。
	localFileName := "C:\\Users\\xy\\Pictures\\test1.jpg"

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com", "", "")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		handleError(err)
	}
}
