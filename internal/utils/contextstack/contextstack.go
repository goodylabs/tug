package contextstack

import "github.com/goodylabs/tug/internal/dto"

type ContextStack struct {
	sshConfig *dto.SSHConfig
	action    string
	resource  string
}

func NewContextStack() *ContextStack {
	return &ContextStack{
		sshConfig: nil,
		action:    "",
		resource:  "",
	}
}

func (c *ContextStack) GetSSHConfig() *dto.SSHConfig {
	return c.sshConfig
}

func (c *ContextStack) SetSSHConfig(sshConfig *dto.SSHConfig) {
	c.sshConfig = sshConfig
}

func (c *ContextStack) ClearSSHConfig() {
	c.sshConfig = nil
}

func (c *ContextStack) GetAction() string {
	return c.action
}

func (c *ContextStack) SetAction(action string) {
	c.action = action
}

func (c *ContextStack) ClearAction() {
	c.action = ""
}

func (c *ContextStack) GetResource() string {
	return c.resource
}

func (c *ContextStack) SetResource(resource string) {
	c.resource = resource
}

func (c *ContextStack) ClearResource() {
	c.resource = ""
}
