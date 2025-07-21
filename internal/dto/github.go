package dto

type GithubReleaseDTO struct {
	Name   string `json:"name"`
	Assets []struct {
		Name        string `json:"name"`
		DownloadUrl string `json:"browser_download_url"`
	} `json:"assets"`
}
