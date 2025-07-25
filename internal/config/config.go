package config

import (
	"log"
	"os"
	"path/filepath"
)

var (
	BASE_DIR   string
	DEVOPS_DIR string
	TUG_ENV    string
)

func LoadProductionConfig() {
	BASE_DIR = getEnvOrError("PWD")
	DEVOPS_DIR = "devops"
}

func LoadDevelopmentConfig() {
	BASE_DIR = filepath.Join(findProjectRoot(), ".development")
	DEVOPS_DIR = "devops"
}

func LoadTestingConfig() {
	BASE_DIR = filepath.Join(findProjectRoot(), ".testing")
	DEVOPS_DIR = "devops"
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
