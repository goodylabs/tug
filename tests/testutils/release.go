package testutils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goodylabs/tug/internal/config"
	"github.com/google/uuid"
)

func CreateTestTugReleaseFile() (string, func()) {
	_ = os.Mkdir(filepath.Join(config.BASE_DIR, ".tmp"), 0755)
	path := filepath.Join(config.BASE_DIR, ".tmp", fmt.Sprintf("release_%s.json", uuid.New().String()))
	return path, func() {
		os.Remove(path)
	}
}
