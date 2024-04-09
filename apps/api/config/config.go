package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
}

// Config func to get env value
func GetString(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
