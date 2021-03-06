package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	Cos struct {
		BucketUrl  string `json:"bucket_url"`
		ServiceUrl string `json:"service_url"`
		SecretId   string `json:"secret_id"`
		SecretKey  string `json:"secret_key"`
	} `json:"cos"`
	Env string `json:"env"`
}

var globalConfig Config

func init() {
	b, err := ioutil.ReadFile(Path)
	if err != nil {
		log.Printf("No config.json? Use config.json.example to create one.")
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
