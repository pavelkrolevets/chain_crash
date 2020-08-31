package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	Addr1 string `yaml:"Address1"`
	PK1 string `yaml:"PK1"`
	Addr2 string `yaml:"Address2"`
	PK2 string `yaml:"PK2"`
	Addr3 string `yaml:"Address3"`
	PK3 string `yaml:"PK3"`
	Addr4 string `yaml:"Address4"`
	PK4 string `yaml:"PK4"`
}

var GlobalConfig Config

func (c *Config) getConfig() *Config  {
	confFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(confFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}


func main(){
	GlobalConfig.getConfig()
	log.Println(GlobalConfig)
}
