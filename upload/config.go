package upload

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Type      string `json:"type"`
	SecretId  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	Region    string `json:"region"`
	Domain    string `json:"domain"`
	Dir       string `json:"dir"`
}

var cnf = new(Config)

func init() {
	bs, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bs, cnf)
	if err != nil {
		panic(err)
	}
}
