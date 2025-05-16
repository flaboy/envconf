package main

import (
	"encoding/json"
	"fmt"

	"github.com/flaboy/envconf"
)

type config struct {
	STR_CFG   string
	INT_CFG   int
	FLOAT_CFG float64

	BOOL_CFG1 bool
	BOOL_CFG2 bool
	BOOL_CFG3 bool
	BOOL_CFG4 bool
	BOOL_CFG5 bool

	CustomVar string            `cfg:"CUSTOM_CFG"`
	JsonVar   map[string]string `cfg:"EXAMPLE_JSON_CFG"`
	JsonVar2  []string          `cfg:"EXAMPLE_JSON_CFG2"`
	Object    ObjectConfig      `cfg:"OBJECT"`
}

type ObjectConfig struct {
	Name string `cfg:"NAME"`
	Age  int    `cfg:"AGE"`
	Addr string `cfg:"ADDR"`
}

var Config *config

func main() {
	Config = &config{}
	fmt.Println(envconf.GetExample(Config))

	err := envconf.Load("env.conf", Config)
	if err != nil {
		panic(err)
	}

	cfgBin, _ := json.MarshalIndent(Config, "", "    ")
	fmt.Printf("%s\n", cfgBin)
}
