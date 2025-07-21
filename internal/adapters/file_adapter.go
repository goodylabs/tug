package adapters

import (
	"os"
	"path/filepath"

	"github.com/goodylabs/docker-swarm-cli/internal/config"

	"bufio"
	"errors"
	"regexp"
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

func GetValueFromShFile(shFilePath string) (string, error) {
	file, err := os.Open(shFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	re := regexp.MustCompile(`^TARGET_IP\s*=\s*["']?([^"']+)["']?$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		matches := re.FindStringSubmatch(line)
		if len(matches) == 2 {
			return matches[1], nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	errorText := "TARGET_IP not found in file:" + shFilePath
	return "", errors.New(errorText)
}
