package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	// load .env file
	err := godotenv.Load(".env")
	fmt.Println("loading .env file")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetString(key string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		panic(fmt.Sprintf("Error loading environment variable \"%s\"\n", key))
	}
	return val
}

func GetInt(key string) uint64 {
	val, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Error loading environment variable \"%s\"\n", key))
	}
	return val
}
