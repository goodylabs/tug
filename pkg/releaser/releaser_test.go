package releaser_test

// func TestPm2LoadStrategy_Execute(t *testing.T) {
// 	testFileName := "ecosystem.config.js"
// 	content := []byte(`module.exports = { deploy: { staging: { user: 'staging-user', host: ['1.2.3.4'] } } };`)

// 	lp := loadproject.NewLoadProject()

// 	_ = os.WriteFile(testFileName, content, 0644)
// 	defer os.Remove(testFileName)

// 	t.Run("successful load via LoadProject factory", func(t *testing.T) {
// 		cfg, err := lp.Execute(loadproject.Pm2Strategy)

// 		if err != nil && strings.Contains(err.Error(), "not found") {
// 			t.Errorf("Error: %s", err.Error())
// 		}

// 		assert.NoError(t, err)
// 		assert.Contains(t, cfg.Con

// func TestReleser_NeedUpdate(t *testing.T) {

// 	file1 := filepath.Join(config.GetBaseDir(), "releaser", "1.json")

// 	assert.Equal(t, "2026-02-26", file1)
// }
