package storage

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Ali struct{}

func (a *Ali) UploadBytes(secretId, secretKey, bucketName, region string, bs []byte, fileName string) (domain string, err error) {
	endpoint := fmt.Sprintf("http://%v.aliyuncs.com", region)

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, secretId, secretKey)
	if err != nil {
		return "", err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	// 上传Byte数组。
	err = bucket.PutObject(fileName, bytes.NewReader(bs))
	if err != nil {
		return "", err
	}
	domain = fmt.Sprintf("http://%v.%v.aliyuncs.com", bucketName, region)
	return domain, nil
}

func (a *Ali) UploadFile(secretId, secretKey, bucketName, region string, filePath, fileName string) {
	endpoint := fmt.Sprintf("http://%v.aliyuncs.com", region)
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, secretId, secretKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 读取本地文件。
	fd, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer fd.Close()

	// 上传文件流。
	err = bucket.PutObject(fileName, fd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func (a *Ali) DeleteObject(secretId, secretKey, bucketName, region, key string) error {
	endpoint := fmt.Sprintf("http://%v.aliyuncs.com", region)

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, secretId, secretKey)
	if err != nil {
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}
	err = bucket.DeleteObject(key)
	return err
}
