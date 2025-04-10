package main

import (
	"encoding/json"
	"fmt"

	"github.com/flaboy/envconf"
)

type StorageConfig struct {
	Type  string       `cfg:"TYPE" default:"local"`
	Local LocalStorage `cfg:"LOCAL"`
	S3    S3Storage    `cfg:"S3"`
}

type LocalStorage struct {
	BasePath string `cfg:"BASE_PATH" default:"storage"`
	BaseURL  string `cfg:"BASE_URL" default:"/storage"`
}

type S3Storage struct {
	AccessKey string `cfg:"ACCESS_KEY"`
	SecretKey string `cfg:"SECRET_KEY"`
	Bucket    string `cfg:"BUCKET"`
	Region    string `cfg:"REGION"`
	Endpoint  string `cfg:"ENDPOINT"`
	PublicURL string `cfg:"PUBLIC_URL"`
}

type Config struct {
	Storage StorageConfig `cfg:"STORAGE"`
}

func main() {
	cfg := &Config{}
	err := envconf.Load("env.conf", cfg)
	if err != nil {
		panic(err)
	}

	cfgBin, _ := json.MarshalIndent(cfg, "", "    ")
	fmt.Printf("%s\n", cfgBin)
}
