package releaser_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/pkg/config"
	"github.com/goodylabs/tug/pkg/releaser"

	"github.com/stretchr/testify/assert"
)

func TestReleaser_CheckIsUpToDate(t *testing.T) {
	baseDir := config.GetBaseDir()
	version := "testing_v1"
	date := "some_date_for_testing"

	t.Run("Valid file with correct Release and LastCheck", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "1.json")
		r := releaser.NewReleaser(path, version, date)
		assert.True(t, r.CheckIsUpToDate())
	})

	t.Run("File with incorrect Release", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "2.json")
		r := releaser.NewReleaser(path, version, date)
		assert.False(t, r.CheckIsUpToDate())
	})

	t.Run("File with incorrect LastCheck", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "3.json")
		r := releaser.NewReleaser(path, version, date)
		assert.False(t, r.CheckIsUpToDate())
	})

	t.Run("Malformed JSON file", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "4.json")
		r := releaser.NewReleaser(path, version, date)
		assert.False(t, r.CheckIsUpToDate())
	})

	t.Run("Empty JSON file", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "5.json")
		r := releaser.NewReleaser(path, version, date)
		assert.False(t, r.CheckIsUpToDate())
	})

	t.Run("Non-existent file", func(t *testing.T) {
		path := filepath.Join(baseDir, "releaser", "non_existent.json")
		r := releaser.NewReleaser(path, version, date)
		assert.False(t, r.CheckIsUpToDate())
	})
}
