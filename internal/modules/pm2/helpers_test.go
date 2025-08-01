package pm2_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/modules/pm2"
	"github.com/stretchr/testify/assert"
)

const output1 = `[{ "name": "pm2-logrotate", "pid": 3105},{  "name": "staging_ro", "pid": 4043331},{ "name": "api-staging", "pid": 4041884}]`

const invalidOutput1 = `>>>> In-memory PM2 is out-of-date, do:
>>>> $ pm2 update
In memory PM2 version: 5.2.2
Local PM2 version: 6.0.8

[{ "name": "pm2-logrotate", "pid": 3105},{  "name": "staging_ro", "pid": 4043331},{ "name": "api-staging", "pid": 4041884}]
`

func TestGetPm2ConfigPath(t *testing.T) {
	t.Run("ecosystem.config.js exists", func(t *testing.T) {
		path, err := pm2.GetPm2ConfigPath(config.BASE_DIR)
		assert.Contains(t, path, "ecosystem.config.js")
		assert.NoError(t, err)
	})

	t.Run("none of ecosystem.config.* exists", func(t *testing.T) {
		path, err := pm2.GetPm2ConfigPath("non-existing-dir")
		assert.Equal(t, "", path)
		assert.ErrorContains(t, err, "file not found")
	})

}
