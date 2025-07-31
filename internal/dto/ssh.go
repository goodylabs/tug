package dto

import "strconv"

type SSHConfig struct {
	User string `json:"user"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (s *SSHConfig) GetSSHConnectionString() string {
	return s.User + "@" + s.Host + ":" + strconv.Itoa(s.Port)
}
