package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariables(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		panic(err)
	}

	return os.Getenv(key)
}
