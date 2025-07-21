package utils

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func UnmarshalLines[T any](input []string) []T {
	var result []T
	for _, line := range input {
		var item T
		line = strings.Trim(line, "'")
		err := json.Unmarshal([]byte(line), &item)
		if err != nil {
			log.Printf("failed to unmarshal line: %s, err: %v", line, err)
			continue
		}
		result = append(result, item)
	}
	return result
}

func UnmarshalJson[T any](input string) T {
	var result T
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		log.Fatal("Can not unmarshal json")
	}
	return result
}

func ReadJSON[T any](path string, v *T) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func WriteJSON[T any](path string, v *T) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
