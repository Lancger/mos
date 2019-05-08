package glo

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// DatabaseConfig Struct
type DatabaseConfig struct {
	Dialect     string `yaml:"dialect"`
	Addr        string `yaml:"addr"`
	AutoMigrate bool   `yaml:"automigrate"`
}

// RedisConfig Struct
type RedisConfig struct {
	MaxIdle        int    `yaml:"max_idle"`
	IdleTimeoutSec int    `yaml:"idle_timeout_sec"`
	Addr           string `yaml:"addr"`
	Password       string `yaml:"password"`
	Expried        int    `yaml:"expried"`
}

// MosAPICfg Struct
type MosAPICfg struct {
	EncryptKey string         `yaml:"encrypt_key"`
	Enable     bool           `yaml:"enable"`
	MaxRequest string         `yaml:"max_request"`
	Deadline   int            `yaml:"deadline"`
	Database   DatabaseConfig `yaml:"database"`
	Redis      RedisConfig    `yaml:"redis"`
	ServerPort int            `yaml:"server_port"`
}

// GlobalConfig Struct
type GlobalConfig struct {
	MosAPI MosAPICfg `yaml:"mos_api"`
}

var (
	// Config Params
	Config GlobalConfig
	// ConfigFile FileName
	ConfigFile string
)

// ParseConfig Function
func ParseConfig(file string) (err error) {
	ConfigFile = file
	var c GlobalConfig
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, &c)
	if err != nil {
		return
	}
	Config = c
	return
}
