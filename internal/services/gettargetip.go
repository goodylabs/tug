package services

import (
	"errors"
	"strings"

	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/utils"
)

func GetTargetIp(scriptAbsPath string) (string, error) {
	fields := []string{
		constants.TARGET_IP_FIELD_LEGACY,
		constants.TARGET_IP_FIELD,
	}

	for _, field := range fields {
		lines, err := utils.GetFileLines(scriptAbsPath)
		if err != nil {
			return "", errors.New("placeholder")
		}

		targetIp := getVariableValueFromLines(lines, field)
		if targetIp != "" {
			return targetIp, nil
		}
	}
	return "", errors.New("placeholder")
}

func getVariableValueFromLines(lines []string, key string) string {
	prefix := key + "="
	for _, line := range lines {
		if after, ok := strings.CutPrefix(line, prefix); ok {
			value := after
			value = strings.Trim(value, `"`)
			return value
		}
	}
	return ""
}
