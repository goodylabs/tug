package pm2

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goodylabs/tug/internal/config"
)

func getPm2ConfigPath() (string, error) {
	options := []string{"ecosystem.config.js", "ecosystem.config.cjs"}
	for _, name := range options {
		path := filepath.Join(config.BASE_DIR, name)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("ecosystem config file not found in %s", config.BASE_DIR)
}

func convertJsFileToJson(path string) error {
	var nodeScript string

	switch {
	case strings.HasSuffix(path, ".js"):
		nodeScript = fmt.Sprintf(ecosystemJsScript, path, tmpJsonPath)
	case strings.HasSuffix(path, ".cjs"):
		nodeScript = fmt.Sprintf(ecosystemCjsScript, path, path, tmpJsonPath)
	default:
		return fmt.Errorf("unsupported ecosystem config file type: %s", path)
	}

	if err := exec.Command("node", "-e", nodeScript).Run(); err != nil {
		return fmt.Errorf("failed to convert ecosystem config file to JSON: %w", err)
	}
	return nil
}
