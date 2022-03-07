package storage

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/taozhang-tt/mdpic/config"
)

type Ali struct {
	SecretId   string
	SecretKey  string
	BucketName string
	Region     string
}

func NewAli(cnf *config.Config) Storage {
	return &Ali{
		SecretId:   cnf.SecretId,
		SecretKey:  cnf.SecretKey,
		BucketName: cnf.Bucket,
		Region:     cnf.Region,
	}
}

func (a *Ali) UploadBytes(fileName string, bs []byte) (domain string, err error) {
	endpoint := fmt.Sprintf("http://%v.aliyuncs.com", a.Region)
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, a.SecretId, a.SecretKey)
	if err != nil {
		return "", err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(a.BucketName)
	if err != nil {
		return "", err
	}

	// 上传Byte数组。
	err = bucket.PutObject(fileName, bytes.NewReader(bs))
	if err != nil {
		return "", err
	}
	domain = fmt.Sprintf("http://%v.%v.aliyuncs.com", a.BucketName, a.Region)
	return domain, nil
}

func (a *Ali) UploadFile(filePath, fileName string) {
	endpoint := fmt.Sprintf("http://%v.aliyuncs.com", a.Region)
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, a.SecretId, a.SecretKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(a.BucketName)
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

func (a *Ali) DeleteObject(key string) error {
	endpoint := fmt.Sprintf("http://%v.aliyuncs.com", a.Region)

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, a.SecretId, a.SecretKey)
	if err != nil {
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(a.BucketName)
	if err != nil {
		return err
	}
	err = bucket.DeleteObject(key)
	return err
}
