package utils

import (
	"encoding/json"
	"log"
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
