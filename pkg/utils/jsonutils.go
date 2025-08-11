package utils

import (
	"encoding/json"
	"os"
)

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
