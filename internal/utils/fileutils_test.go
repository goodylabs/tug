package utils_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetFileLinesOk(t *testing.T) {
	tests := []struct {
		envDir      string
		linesNumber int
	}{
		{"localhost", 11},
		{"production", 10},
	}

	for _, tt := range tests {
		scriptAbsPath := filepath.Join(config.BASE_DIR, config.DEVOPS_DIR, tt.envDir, "deploy.sh")
		lines, err := utils.GetFileLines(scriptAbsPath)
		assert.Equal(t, tt.linesNumber, len(lines))
		assert.NoError(t, err)
	}
}
