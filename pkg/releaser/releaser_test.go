package releaser_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/pkg/config"
	"github.com/goodylabs/tug/pkg/releaser"

	"github.com/stretchr/testify/assert"
)

func TestReleaser_CheckNeedUpdate(t *testing.T) {
	tests := []struct {
		name               string
		filePath           string
		expectedNeedUpdate bool
	}{
		{
			name:               "Valid file with correct Release and LastCheck",
			filePath:           filepath.Join(config.GetBaseDir(), "releaser", "1.json"),
			expectedNeedUpdate: true,
		},
		{
			name:               "File with incorrect Release",
			filePath:           filepath.Join(config.GetBaseDir(), "releaser", "2.json"),
			expectedNeedUpdate: false,
		},
		{
			name:               "File with incorrect LastCheck",
			filePath:           filepath.Join(config.GetBaseDir(), "releaser", "3.json"),
			expectedNeedUpdate: false,
		},
		{
			name:               "Malformed JSON file",
			filePath:           filepath.Join(config.GetBaseDir(), "releaser", "4.json"),
			expectedNeedUpdate: false,
		},
		{
			name:               "Empty JSON file",
			filePath:           filepath.Join(config.GetBaseDir(), "releaser", "5.json"),
			expectedNeedUpdate: false,
		},
		{
			name:               "Non-existent file",
			filePath:           filepath.Join(config.GetBaseDir(), "releaser", "non_existent.json"),
			expectedNeedUpdate: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := releaser.NewReleaser(tt.filePath)
			assert.Equal(t, tt.expectedNeedUpdate, r.CheckNeedUpdate())
		})
	}
}
