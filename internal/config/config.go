package config

import (
	"log"
	"os"
	"path/filepath"
)

var (
	BASE_DIR   string
	USE_MOCKS  bool
	DEVOPS_DIR string
)

var DOCKER_HOST_ENV string

func init() {
	BASE_DIR = findProjectRoot()
	USE_MOCKS = bool(getEnvOrDefault("USE_MOCKS", "") == "true")
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
