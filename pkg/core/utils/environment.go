package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {
  godotenv.Load(".env")
  return os.Getenv(key)
}