package config

import (
	"log"
	"os"
	"path/filepath"
)

var (
	baseDir       string
	homeDir       string
	tugConfigPath string
)

func GetBaseDir() string {
	if baseDir == "" {
		tugEnv := os.Getenv("TUG_ENV")
		if tugEnv == "development" || tugEnv == "testing" {
			projectRoot := findProjectRoot()
			baseDir = filepath.Join(projectRoot, "."+tugEnv)
		} else {
			baseDir = getEnvOrError("PWD")
		}
	}
	return baseDir
}

func GetHomeDir() string {
	if homeDir == "" {
		tugEnv := os.Getenv("TUG_ENV")
		if tugEnv == "development" || tugEnv == "testing" {
			projectRoot := findProjectRoot()
			homeDir = filepath.Join(projectRoot, "."+tugEnv)
		} else {
			homeDir = getEnvOrError("HOME")
		}
	}
	return homeDir
}

func GetTugConfigPath() string {
	if tugConfigPath == "" {
		filepath.Join(GetBaseDir(), "ecosystem.config.js")
	}
	return tugConfigPath
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
