package adapters

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goodylabs/docker-swarm-cli/internal/config"

	"strings"
)

func ListDirectories(path string) ([]string, error) {
	absPath := filepath.Join(config.BASE_DIR, path)
	entries, err := os.ReadDir(absPath)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}
	return dirs, nil
}

func GetValueFromShFile(shFilePath string, key string) (string, error) {
	content, err := os.ReadFile(shFilePath)
	if err != nil {
		return "", err
	}

	for line := range strings.SplitSeq(string(content), "\n") {
		line = strings.TrimSpace(line)
		if after, ok := strings.CutPrefix(line, key+"="); ok {
			value := after
			value = strings.Trim(value, `"'`)
			return value, nil
		}
	}

	return "", fmt.Errorf("%s not found in file: %s", key, shFilePath)
}
