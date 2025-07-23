package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	BASE_DIR   string
	DEVOPS_DIR string
)

func init() {
	godotenv.Load(".env")
	BASE_DIR = findProjectRoot()
	DEVOPS_DIR = getEnvOrDefault("DEVOPS_DIR", "devops")
}

func getEnvOrDefault(envName string, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	return value
}

func findProjectRoot() string {
	dir := getEnvOrDefault("BASE_DIR", "")
	if dir != "" {
		abs, err := filepath.Abs(dir)
		if err != nil {
			log.Fatal(err)
		}
		return abs
	}
	abs, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}
	return abs
}
