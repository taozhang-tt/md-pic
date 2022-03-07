package manage

import (
	"fmt"
	"time"

	"golang.design/x/clipboard"

	"github.com/taozhang-tt/mdpic/config"
	"github.com/taozhang-tt/mdpic/storage"
)

func UploadFromClipboard() {
	cnf := config.GetConfig()

	bs := clipboard.Read(clipboard.FmtImage)
	if len(bs) == 0 {
		fmt.Printf("clipboard empty\n")
		return
	}
	fileName := fmt.Sprintf("%v.png", time.Now().Unix())
	if cnf.Dir != "" {
		fileName = cnf.Dir + "/" + fileName
	}
	client := storage.CreateStorage(cnf)
	domain, err := client.UploadBytes(fileName, bs)
	if err != nil {
		fmt.Printf("upload err: %v\n", err)
	}
	if cnf.Domain != "" {
		domain = cnf.Domain
	}
	addr := fmt.Sprintf("%v/%v", domain, fileName)

	_ = clipboard.Write(clipboard.FmtText, []byte(addr))
	fmt.Printf("pic upload success\n")
}
func DeleteObject(keys []string) {
	cnf := config.GetConfig()
	client := storage.CreateStorage(cnf)

	for _, key := range keys {
		fileName := fmt.Sprintf("%v.png", key)
		if cnf.Dir != "" {
			fileName = cnf.Dir + "/" + fileName
		}
		err := client.DeleteObject(fileName)
		if err != nil {
			fmt.Printf("Delete fail: %v\n", err)
		} else {
			fmt.Printf("Delete %v success\n", key)
		}
	}
}
