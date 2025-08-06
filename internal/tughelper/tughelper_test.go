package tughelper_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/tughelper"
	"github.com/stretchr/testify/assert"
)

type tugHelperTestCase struct {
	promptChoices []int
	expected      string
	value         string
}

func TestSetGetTugConfig(t *testing.T) {
	os.Remove(config.TUG_CONFIG_PATH)

	_, err := tughelper.GetTugConfig()
	assert.Error(t, err)

	tugConfig := &tughelper.TugConfig{
		SSHKeyPath: "test_key_path",
	}
	err = tughelper.SetTugConfig(tugConfig)

	assert.NoError(t, err, "Expected no error, got: %v", err)
	retrievedConfig, err := tughelper.GetTugConfig()

	assert.NoError(t, err, "Expected no error, got: %v", err)
	assert.Equal(t, tugConfig.SSHKeyPath, retrievedConfig.SSHKeyPath)

}

func TestGetAvailableSSHFilesOK(t *testing.T) {
	sshDirPath := filepath.Join(config.BASE_DIR, ".ssh")
	sshFiles, err := tughelper.GetAvailableSSHFiles(sshDirPath)
	assert.NoError(t, err, "Expected no error, got: %v", err)
	assert.True(t, len(sshFiles) == 2)
	assert.Contains(t, sshFiles[0], "id_rsa")
	assert.Contains(t, sshFiles[1], "some_priv_key")
}
