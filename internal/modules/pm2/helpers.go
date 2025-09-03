package pm2

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goodylabs/tug/pkg/config"
)

func GetPm2ConfigPath(dir string) (string, error) {
	options := []string{"ecosystem.config.cjs", "ecosystem.config.js"}
	for _, name := range options {
		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("ecosystem.config.cjs/js file not found in %s", config.GetBaseDir())
}

func ConvertJsFileToJson(path string) error {
	var nodeScript string

	switch {
	case strings.HasSuffix(path, ".js"):
		nodeScript = fmt.Sprintf(ecosystemJsScript, path, tmpJsonPath)
	case strings.HasSuffix(path, ".cjs"):
		nodeScript = fmt.Sprintf(ecosystemCjsScript, path, path, tmpJsonPath)
	default:
		return fmt.Errorf("unsupported ecosystem.config.cjs/js file type: %s", path)
	}

	if err := exec.Command("node", "-e", nodeScript).Run(); err != nil {
		return fmt.Errorf("failed to convert ecosystem.config.cjs/js file to JSON: %w", err)
	}
	return nil
}
