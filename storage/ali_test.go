package storage

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAliUploadBytes(t *testing.T) {
	Convey("AliUploadBytes", t, func() {
		secretId := ""
		secretKey := ""
		bucket := ""
		region := ""

		ali := new(Ali)
		ali.UploadBytes(secretId, secretKey, bucket, region, []byte("Hello World!"), "uploadbytes.txt")
	})
}

func TestAliUploadFile(t *testing.T) {
	Convey("AliUploadFile", t, func() {
		secretId := ""
		secretKey := ""
		bucket := ""
		region := ""

		ali := new(Ali)
		ali.UploadFile(secretId, secretKey, bucket, region, "/Users/tt/Code/go-code/md-pic/README.md", "readme.md")
	})
}

func TestAliDeleteFile(t *testing.T) {
	Convey("AliUploadFile", t, func() {
		secretId := ""
		secretKey := ""
		bucket := ""
		region := ""
		key := "note/1638360157.png"

		ali := new(Ali)
		err := ali.DeleteObject(secretId, secretKey, bucket, region, key)
		fmt.Println(err)
	})
}
