package initialize

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}

}
