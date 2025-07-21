package utils

import (
	"bufio"
	"os"
)

// func ListDirectories(absPath string) ([]string, error) {
// 	entries, err := os.ReadDir(absPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var dirs []string
// 	for _, entry := range entries {
// 		if entry.IsDir() {
// 			dirs = append(dirs, entry.Name())
// 		}
// 	}
// 	return dirs, nil
// }

func GetFileLines(scriptAbsPath string) ([]string, error) {
	file, err := os.Open(scriptAbsPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
