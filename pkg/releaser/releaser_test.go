package releaser

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/pkg/config"

	"github.com/stretchr/testify/assert"
)

func TestReleaser_checkIsUpToDate(t *testing.T) {
	baseDir := config.GetBaseDir()
	version := "testing_release"
	date := "testing_lastCheck"

	t.Run("return true on correct values", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "1.json")
		r := NewReleaser(path, version, date)
		assert.True(t, r.checkIsUpToDate())
	})

	t.Run("return false on incorrect release", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "2.json")
		r := NewReleaser(path, version, date)
		assert.False(t, r.checkIsUpToDate())
	})

	t.Run("return false on incorrect last check", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "3.json")
		r := NewReleaser(path, version, date)
		assert.False(t, r.checkIsUpToDate())
	})

	t.Run("return false on non JSON format", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "4.json")
		r := NewReleaser(path, version, date)
		assert.False(t, r.checkIsUpToDate())
	})

	t.Run("return false on empty JSON file", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "5.json")
		r := NewReleaser(path, version, date)
		assert.False(t, r.checkIsUpToDate())
	})

	t.Run("return false on empty file", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "6.json")
		r := NewReleaser(path, version, date)
		assert.False(t, r.checkIsUpToDate())
	})

	t.Run("return false on non-existent file", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "non_existent.json")
		r := NewReleaser(path, version, date)
		assert.False(t, r.checkIsUpToDate())
	})
}
