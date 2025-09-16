package services

import "github.com/goodylabs/tug/internal/modules"

const continueMsg = "echo 'Done, press Enter to continue...' && read"

func GetActionTemplates() []modules.TechCmdTemplate {
	return []modules.TechCmdTemplate{
		{
			Display:  "[bash] bash",
			Template: "bash",
		},
		{
			Display:  "[bash] htop",
			Template: "htop",
		},
		{
			Display:  "[bash] btop",
			Template: "btop",
		},
		{
			Display:  "[bash] df -h",
			Template: "df -h && " + continueMsg,
		},
		{
			Display:  "[bash] free -h",
			Template: "free -h && " + continueMsg,
		},
	}
}
