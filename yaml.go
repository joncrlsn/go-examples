package main

import (
	"bytes"
	"github.com/spf13/viper"
	"log"
)

func main() {
	//
	// Initialize from our config file
	//
	//viper.SetConfigType("yaml")
	//viper.SetConfigName("yaml-config") // name of config file (without extension)
	//viper.AddConfigPath(".")
	//viper.WatchConfig()
	//err := viper.ReadInConfig() // Find and read the config file
	//if err != nil {
	//log.Fatalln(err)
	//}

	// Parse bytes
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
Hacker: false
`)
	yaml := viper.New()
	yaml.SetConfigType("yaml")
	yaml.ReadConfig(bytes.NewBuffer(yamlExample))

	log.Println("clothing.jacket", yaml.Get("clothing.jacket"))
	log.Println("age", yaml.Get("age"))
	log.Println("hobbies", yaml.Get("hobbies"))
	log.Println("hacker", yaml.Get("Hacker"))
}
