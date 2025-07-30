package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	BASE_DIR        string
	HOME_DIR        string
	TUG_CONFIG_PATH string
	PM2_CONFIG_PATH string
)

func Load() {
	godotenv.Load(".env")
	tugEnv := os.Getenv("TUG_ENV")
	fmt.Println("Using TUG_ENV:", tugEnv)
	projectRoot := findProjectRoot()

	switch tugEnv {
	case "development":
		BASE_DIR = filepath.Join(projectRoot, ".development")
		HOME_DIR = getEnvOrError("HOME")
	case "testing":
		BASE_DIR = filepath.Join(projectRoot, ".testing")
		HOME_DIR = filepath.Join(projectRoot, ".testing")
	default:
		BASE_DIR = getEnvOrError("PWD")
		HOME_DIR = getEnvOrError("HOME")
	}

	TUG_CONFIG_PATH = filepath.Join(HOME_DIR, ".tug", "tugconfig.json")
	PM2_CONFIG_PATH = filepath.Join(BASE_DIR, "ecosystem.config.js")
}

func loadTestingConfig() {
	projectRoot := findProjectRoot()
	BASE_DIR = filepath.Join(projectRoot, ".testing")
	HOME_DIR = filepath.Join(projectRoot, ".testing")
	TUG_CONFIG_PATH = filepath.Join(HOME_DIR, ".tug", "tugconfig.json")
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
