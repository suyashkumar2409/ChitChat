package config

import (
	"log"
	"os"
	"encoding/json"
)

type Config struct{
	Address string
	ReadTimeout int64
	WriteTimeout int64
	Static string
	Version string
}

var config Config
var logger *log.Logger

func init(){
	loadConfig()
	loadLogger()
}

func loadConfig(){
	file, err := os.Open("config/config.json")
	if err != nil{
		log.Fatalln("Failed to open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Config{}
	err = decoder.Decode(&config)
	if err != nil{
		log.Fatalln("Unable to get configuration from file", err)
	}
}

func loadLogger(){
	file, err := os.OpenFile("chitchat.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil{
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate | log.Ltime | log.Lshortfile)
}

func GetAddress() string{
	return config.Address
}

func GetReadTimeout() int64{
	return config.ReadTimeout
}

func GetWriteTimeout() int64{
	return config.WriteTimeout
}

func GetStatic() string{
	return config.Static
}

func GetVersion() string{
	return config.Version
}

func Info(args ...interface{}){
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func Error(args ...interface{}){
	logger.SetPrefix("Error ")
	logger.Println(args...)
}

func Warning(args ...interface{}){
	logger.SetPrefix("Warning ")
	logger.Println(args...)
}

