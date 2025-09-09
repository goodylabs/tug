package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYAML[T any](path string, v *T) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}

func WriteYAML[T any](path string, v *T) error {
	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
