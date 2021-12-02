package manage

import (
	"fmt"
	"time"

	"golang.design/x/clipboard"

	"mdpic/storage"
)

func UploadFromClipboard() {
	bs := clipboard.Read(clipboard.FmtImage)
	if len(bs) == 0 {
		fmt.Printf("clipboard empty\n")
		return
	}
	fileName := fmt.Sprintf("%v.png", time.Now().Unix())
	if cnf.Dir != "" {
		fileName = cnf.Dir + "/" + fileName
	}
	ali := new(storage.Ali)
	domain, err := ali.UploadBytes(cnf.SecretId, cnf.SecretKey, cnf.Bucket, cnf.Region, bs, fileName)
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
	for _, key := range keys {
		fileName := fmt.Sprintf("%v.png", key)
		if cnf.Dir != "" {
			fileName = cnf.Dir + "/" + fileName
		}
		ali := new(storage.Ali)
		err := ali.DeleteObject(cnf.SecretId, cnf.SecretKey, cnf.Bucket, cnf.Region, fileName)
		if err != nil {
			fmt.Printf("Delete fail: %v\n", err)
		} else {
			fmt.Printf("Delete %v success\n", key)
		}
	}
}
