package storage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAliUploadBytes(t *testing.T) {
	Convey("AliUploadBytes", t, func() {
		secretId := "LTAI4G5aNwjYSLrqrmaDzFMq"
		secretKey := "gONBIxAdB9KuQULq769bveRsrRmk3t"
		bucket := "wuchang-tt"
		region := "oss-cn-hongkong"

		ali := new(Ali)
		ali.UploadBytes(secretId, secretKey, bucket, region, []byte("Hello World!"), "uploadbytes.txt")
	})
}

func TestAliUploadFile(t *testing.T) {
	Convey("AliUploadFile", t, func() {
		secretId := "LTAI4G5aNwjYSLrqrmaDzFMq"
		secretKey := "gONBIxAdB9KuQULq769bveRsrRmk3t"
		bucket := "wuchang-tt"
		region := "oss-cn-hongkong"

		ali := new(Ali)
		ali.UploadFile(secretId, secretKey, bucket, region, "/Users/tt/Code/go-code/md-pic/README.md", "readme.md")
	})
}
