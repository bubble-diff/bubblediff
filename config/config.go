package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	Path        = "config.json"
	IdentityKey = "userid"
)

type Config struct {
	ListenAddr   string `json:"listen_addr"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	// BubblediffUi 用于允许跨域访问
	BubblediffUi string `json:"bubblediff_ui"`
	Mongo        struct {
		Url         string `json:"url"`
		Collections struct {
			User string `json:"user"`
			Task string `json:"task"`
		} `json:"collections"`
	} `json:"mongo"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
	} `json:"redis"`
	Env string `json:"env"`
}

var globalConfig Config

func init() {
	b, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &globalConfig)
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	return &globalConfig
}
