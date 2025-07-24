package tughelper_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/tughelper"
	"github.com/goodylabs/tug/tests/testutils"
	"github.com/stretchr/testify/assert"
)

func TestGetReleaseWhenEmptyFile(t *testing.T) {

	path, cleanup := testutils.CreateTestTugReleaseFile()
	defer cleanup()

	_, err := tughelper.GetTugRelease(path)
	assert.Error(t, err)

}

func TestCreateDefaultReleaseFile(t *testing.T) {

	path, cleanup := testutils.CreateTestTugReleaseFile()
	defer cleanup()

	err := tughelper.CreateDefaultTugRelease(path)
	assert.NoError(t, err)

	tugRelease, err := tughelper.GetTugRelease(path)
	assert.NoError(t, err)
	assert.Equal(t, tugRelease.CurrentVersion, "N/A")
	assert.Equal(t, tugRelease.LastVersionCheckDate, "N/A")

}

func TestLastVersionCheckDate(t *testing.T) {

	path, cleanup := testutils.CreateTestTugReleaseFile()
	defer cleanup()

	err := tughelper.CreateDefaultTugRelease(path)
	assert.NoError(t, err)

	tugRelease, err := tughelper.GetTugRelease(path)
	assert.NoError(t, err)
	assert.Equal(t, tugRelease.CurrentVersion, "N/A")
	assert.Equal(t, tugRelease.LastVersionCheckDate, "N/A")

}

func TestNoNeedToChangeRunCurl(t *testing.T) {

	path, cleanup := testutils.CreateTestTugReleaseFile()
	defer cleanup()

	err := tughelper.CreateDefaultTugRelease(path)
	assert.NoError(t, err)

	tugRelease, err := tughelper.GetTugRelease(path)
	assert.NoError(t, err)
	assert.Equal(t, tugRelease.CurrentVersion, "N/A")
	assert.Equal(t, tugRelease.LastVersionCheckDate, "N/A")

}
