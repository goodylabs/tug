package docker

import (
	"errors"
	"strings"

	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/utils"
)

func (d *DockerManager) GetTargetIpFromScript(scriptPath string) (string, error) {
	for _, field := range []string{constants.TARGET_IP_FIELD_LEGACY, constants.TARGET_IP_FIELD} {
		lines, err := utils.GetFileLines(scriptPath)
		if err != nil {
			return "", err
		}

		if ip := extractVariable(lines, field); ip != "" {
			return ip, nil
		}
	}
	return "", errors.New("could not find TARGET_IP in deploy.sh")
}

func extractVariable(lines []string, key string) string {
	prefix := key + "="
	for _, line := range lines {
		if value, ok := strings.CutPrefix(line, prefix); ok {
			return strings.Trim(value, `"`)
		}
	}
	return ""
}
