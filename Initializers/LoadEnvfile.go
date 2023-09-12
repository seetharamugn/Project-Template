package Initializers

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf("error loading .env file")
	}
}
