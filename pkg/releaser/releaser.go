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

func NewReleaser(releaserFilePath string) *Releaser {
	return &Releaser{
		releaserFilePath: releaserFilePath,
	}
}

type releaserFile struct {
	Release   string `json:"release"`
	LastCheck string `json:"lastCheck"`
}

func (r *Releaser) CheckNeedUpdate() bool {
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

	if rf.Release != "v1.31" {
		return false
	}

	if rf.LastCheck != "1-1-1" {
		return false
	}

	return true
}
