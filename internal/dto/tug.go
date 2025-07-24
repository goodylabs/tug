package dto

const (
	PlatformLinux  string = "linux"
	PlatformDarwin string = "darwin"
)

type TugReleaseDTO struct {
	Platform             string `json:"platform"`
	CurrentVersion       string `json:"currentVersion"`
	LastVersionCheckDate string `json:"LastVersionCheckDate"`
}
