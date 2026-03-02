package loadproject_test

import (
	"os"
	"strings"
	"testing"

	"github.com/goodylabs/tug/internal/modules/loadproject"
	"github.com/stretchr/testify/assert"
)

func TestPm2LoadStrategy_Execute(t *testing.T) {
	testFileName := "ecosystem.config.js"
	content := []byte(`module.exports = { deploy: { staging: { user: 'staging-user', host: ['1.2.3.4'] } } };`)

	lp := loadproject.NewLoadProject()

	_ = os.WriteFile(testFileName, content, 0644)
	defer os.Remove(testFileName)

	t.Run("successful load via LoadProject factory", func(t *testing.T) {
		cfg, err := lp.Execute(loadproject.Pm2Strategy)

		if err != nil && strings.Contains(err.Error(), "not found") {
			t.Errorf("Error: %s", err.Error())
		}

		assert.NoError(t, err)
		assert.Contains(t, cfg.Config, "staging")
		assert.Equal(t, "staging-user", cfg.Config["staging"].User)
		assert.Equal(t, []string{"xxx.xxx.xxx.xxx"}, cfg.Config["staging"].Hosts)
	})
}
