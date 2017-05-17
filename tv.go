package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

// Config for reverse proxy settings and RBAC users and groups
// Unmarshalled from config on disk
type Config struct {
	ListenInterface string
	ListenPort      int
	Thing           map[string]SomeOtherThing
}

type SomeOtherThing struct {
	Stuff string
}

func (C *Config) getConf() *Config {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, C)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return C
}

func main() {
	// Get config
	viper.SetConfigName("config") // name of config file (without extension)
	// yaml, toml, json, ini, it don't care
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	var C Config
	err = viper.Unmarshal(&C)
	if err != nil {
		panic(fmt.Errorf("Unable to decode config into struct: %v", err))
	}

	fmt.Println("Output from viper:")
	spew.Dump(C)

	var D Config
	D.getConf()

	fmt.Println("Output from go-yaml:")
	spew.Dump(D)
}
