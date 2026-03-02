package config

import (
	"log"
	"os"
	"path/filepath"
)

const (
	ModeDev  = "development"
	ModeTest = "testing"
	ModeProd = "production"
)

var TugEnv string

var (
	baseDir string
	homeDir string
)

func init() {
	env := os.Getenv("TUG_ENV")
	if env == "" {
		TugEnv = ModeProd
		return
	}
	switch env {
	case ModeDev, ModeTest, ModeProd:
		TugEnv = env
	default:
		log.Fatalf("\n[FATAL] Invalid TUG_ENV value: '%s'. Allowed: %s, %s, %s.",
			env, ModeDev, ModeTest, ModeProd)
	}
}

func GetMode() string {
	return TugEnv
}

func GetBaseDir() string {
	if baseDir == "" {
		mode := GetMode()
		if mode == ModeDev || mode == ModeTest {
			projectRoot := findProjectRoot()
			baseDir = filepath.Join(projectRoot, "."+mode)
		} else {
			baseDir = getEnvOrError("PWD")
		}
	}
	return baseDir
}

func GetHomeDir() string {
	if homeDir == "" {
		mode := GetMode()
		if mode == ModeTest {
			projectRoot := findProjectRoot()
			homeDir = filepath.Join(projectRoot, "."+mode)
		} else {
			homeDir = getEnvOrError("HOME")
		}
	}
	return homeDir
}

func getEnvOrError(envName string) string {
	value := os.Getenv(envName)
	if value == "" {
		if envName == "PWD" {
			dir, _ := os.Getwd()
			return dir
		}
		log.Fatalf("[FATAL] Environment variable %s is not set.", envName)
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
			log.Fatal("[FATAL] Could not find project root (go.mod).")
		}
		dir = parent
	}
}
