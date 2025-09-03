package services

import (
	"bufio"
	"os"
	"strings"

	"github.com/goodylabs/tug/pkg/utils"
)

func ListEnvs(devopsDirPath string) ([]string, error) {
	return utils.ListDirsOnPath(devopsDirPath)
}

func getShellVariableValue(scriptPath, variable string) (string, error) {
	file, err := os.Open(scriptPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.HasPrefix(line, variable+"=") {
			continue
		}

		rightPart := strings.TrimPrefix(line, variable+"=")

		var clearedStr = rightPart
		clearedStr = strings.ReplaceAll(clearedStr, `"`, "")
		clearedStr = strings.ReplaceAll(clearedStr, `'`, "")
		clearedStr = strings.ReplaceAll(clearedStr, "`", "")

		return clearedStr, nil
	}

	return "", nil
}

func GetSingleIpFromShellFile(scriptPath, variable string) string {
	value, _ := getShellVariableValue(scriptPath, variable)
	return value
}

func GetMultipleIpsFromShellScript(scriptPath, variable string) []string {
	value, _ := getShellVariableValue(scriptPath, variable)

	if value == "" {
		return []string{}
	}

	var trimmedValue = value
	trimmedValue = strings.TrimPrefix(trimmedValue, "(")
	trimmedValue = strings.TrimSuffix(trimmedValue, ")")

	return strings.Split(trimmedValue, " ")
}
