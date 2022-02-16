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
	WebAddress   string `json:"web_address"`
	MongoUrl     string `json:"mongo_url"`
	Redis        struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
	} `json:"redis"`
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
