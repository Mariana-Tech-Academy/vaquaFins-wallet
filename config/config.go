package config

import (
	"github.com/joho/godotenv"
	"log"
)

//This function loads enviromental variables from .env file

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

}
