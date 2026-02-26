package releaser

import (
	"encoding/json"
	"os"
)

type Releaser struct {
	version          string
	todaysDate       string
	releaserFilePath string
}

func NewReleaser(releaserFilePath, version, todaysDate string) *Releaser {
	return &Releaser{
		releaserFilePath: releaserFilePath,
		version:          version,
		todaysDate:       todaysDate,
	}
}

type releaserFile struct {
	Version   string `json:"version"`
	LastCheck string `json:"lastCheck"`
}

func (r *Releaser) CheckIsUpToDate() bool {
	if _, err := os.Stat(r.releaserFilePath); os.IsNotExist(err) {
		return false
	}

	fileContent, err := os.ReadFile(r.releaserFilePath)
	if err != nil {
		return false
	}

	var rf releaserFile
	if err := json.Unmarshal(fileContent, &rf); err != nil {
		return false
	}

	if rf.Version != r.version {
		return false
	}

	if rf.LastCheck != r.todaysDate {
		return false
	}

	return true
}
