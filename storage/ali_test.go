package storage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/taozhang-tt/mdpic/config"
)

func TestAliUploadBytes(t *testing.T) {
	Convey("AliUploadBytes", t, func() {
		ali := NewAli(&config.Config{})
		ali.UploadBytes("uploadbytes.txt", []byte("Hello World!"))
	})
}
