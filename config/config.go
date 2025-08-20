package config

import (
	"github.com/joho/godotenv"
	"log"
)

//This function loads environmental variables from .env file

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

}
