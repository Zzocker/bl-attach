package main

import (
	"os"
	"gopkg.in/yaml.v2"
)

type config struct {
	Port int `yaml:"port"`
}


func readConfig(path string) (*config,error){
	file,err := os.Open(path)
	if err!=nil{
		return nil,err
	}
	defer file.Close()
	var cfg config
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err!=nil{
		return nil,err
	}
	return &cfg,nil
}