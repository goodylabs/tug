package pm2_test

// func TestGetPm2ProcessesOk(t *testing.T) {
// 	choices := []int{1, 2}
// 	expected := []string{"api-staging", "pm2-logrotate"}

// 	for id := range expected {
// 		pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{choices[id]}, output1, nil)
// 		resources, err := pm2Manager.GetAvailableResources(&dto.SSHConfig{})
// 		assert.NoError(t, err)
// 		assert.Equal(t, expected[id], resources)
// 	}

// }

// func TestGetPm2ProcessesCommandError(t *testing.T) {
// 	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{0}, "", errors.New("command error"))
// 	resource, err := pm2Manager.GetAvailableResources(&dto.SSHConfig{})
// 	assert.ErrorContains(t, err, "running PM2 jlist command: command error")
// 	assert.Equal(t, "", resource)
// }

// func TestGetPm2ProcessesInvalidOutput(t *testing.T) {
// 	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{0}, invalidOutput1, nil)
// 	resource, err := pm2Manager.GetAvailableResources(&dto.SSHConfig{})
// 	assert.ErrorContains(t, err, "parsing PM2 list output")
// 	assert.Equal(t, "", resource)
// }
