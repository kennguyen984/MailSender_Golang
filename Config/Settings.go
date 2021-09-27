package Config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvLoader() map[string]string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	keys := []string{
		"EMAIL_API_KEY",
		"EMAIL_HOST",
		"EMAIL_PORT",
		"EMAIL_AUTH_USER",
		"EMAIL_AUTH_PASSWORD",
		"EMAIL_TYPE",
	}
	vars := map[string]string{}
	for _, item := range keys {
		vars[item] = os.Getenv(item)
	}
	return vars
}
