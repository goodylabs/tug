package dto

type SSHConfigDTO struct {
	User string `json:"user"`
	Host string `json:"host"`
	Port int    `json:"port"`
}
