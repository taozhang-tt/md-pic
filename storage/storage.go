package storage

import "github.com/taozhang-tt/mdpic/config"

type Storage interface {
	UploadBytes(path string, bs []byte) (string, error)
	DeleteObject(key string) error
}

var storages map[string]func(config *config.Config) Storage

func init() {
	storages = make(map[string]func(config *config.Config) Storage)
	storages["github"] = NewGithub
	storages["ali"] = NewAli
}

func CreateStorage(cnf *config.Config) Storage {
	f, ok := storages[cnf.Type]
	if ok {
		return f(cnf)
	}
	panic("type not fund")
}
