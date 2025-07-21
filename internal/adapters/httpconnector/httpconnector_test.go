package httpconnector_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/adapters/httpconnector"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/stretchr/testify/assert"
)

var connector ports.HttpConnector

func init() {
	connector = httpconnector.NewHttpConnector()
}

func TestHttpGetForGithubDto(t *testing.T) {
	var release dto.GithubReleaseDTO
	err := connector.HttpGetReq(constants.GITHUB_LAST_RELEASE, &release)
	assert.NoError(t, err)
}
