package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Type        string `json:"type"`
	SecretId    string `json:"secret_id"`
	SecretKey   string `json:"secret_key"`
	Bucket      string `json:"bucket"`
	Region      string `json:"region"`
	Domain      string `json:"domain"`
	Dir         string `json:"dir"`
	GithubToken string `json:"github_token"`
	GithubOwner string `json:"github_owner"`
	GithubRepo  string `json:"github_repo"`
}

var cnf = new(Config)

func GetConfig() *Config {
	return cnf
}

func init() {
	bs, err := ioutil.ReadFile("/etc/mdpic.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bs, cnf)
	if err != nil {
		panic(err)
	}
}
