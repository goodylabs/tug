package pm2

import (
	"encoding/json"
)

const (
	JLIST_CMD = `pm2 jlist | sed -n '/^\[/,$p'`
)

func (p *Pm2Manager) JsonOutputHandler(output string, dtoStruct any) error {
	return json.Unmarshal([]byte(output), &dtoStruct)
}
