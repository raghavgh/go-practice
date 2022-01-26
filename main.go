package main

import (
	"encoding/json"
	"os"
)

var config Config

type Config struct {
	Database struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		DBName   string `json:"dbname"`
		Password string `json:"password"`
		Port     string `json:"port"`
	} `json:"databaseConfig"`
}

func LoadConfig(file string) Config {
	var configRes Config
	confFile, err := os.Open(file)
	defer confFile.Close()
	if err != nil {
		panic("error in opening conf file")
	}
	jsonParser := json.NewDecoder(confFile)
	jsonParser.Decode(&configRes)
	return configRes
}

func main() {
	config = LoadConfig("resources/config.json")
	a := App{}
	db := config.Database
	a.Initialize(db.User, db.Password, db.DBName)
	a.Run(":8010")
}
