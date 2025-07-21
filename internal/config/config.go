package config

import (
	"log"
	"os"
	"path/filepath"
)

var (
	BASE_DIR   string
	DEVOPS_DIR string
	TESTING    bool
)

const (
	DOCKER_API_VERSION = "1.41"
	LOCAL_DOCKER_HOST  = "unix:///var/run/docker.sock"
)

func init() {
	BASE_DIR = findProjectRoot()
	DEVOPS_DIR = getEnvOrDefault("DEVOPS_DIR", ".example-dir")
	TESTING = bool(getEnvOrDefault("TESTING", "") == "true")
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
