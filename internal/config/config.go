package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	BASE_DIR string
)

func init() {
	godotenv.Load(".env")
	tugEnv := os.Getenv("TUG_ENV")

	switch tugEnv {
	case "development":
		loadDevelopmentConfig()
	case "testing":
		loadTestingConfig()
	default:
		loadProductionConfig()
	}
}

func loadProductionConfig() {
	BASE_DIR = getEnvOrError("PWD")
}

func loadDevelopmentConfig() {
	BASE_DIR = filepath.Join(findProjectRoot(), ".development")
}

func loadTestingConfig() {
	BASE_DIR = filepath.Join(findProjectRoot(), ".testing")
}

func getEnvOrError(envName string) string {
	value := os.Getenv(envName)
	if value == "" {
		log.Fatalf("Environment variable %s is not set - tug does not support your shell configuration...", envName)
	}
	return value
}

func findProjectRoot() string {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			log.Fatal("Could not find project root with go.mod file")
		}
		dir = parent
	}
}
