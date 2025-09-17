package services

const continueMsg = "echo 'Done, press Enter to continue...' && read"

func GetActionTemplates() map[string]string {
	return map[string]string{
		"bash  --  bash":    "bash",
		"bash  --  htop":    "htop",
		"bash  --  btop":    "btop",
		"bash  --  df -h":   "df -h && " + continueMsg,
		"bash  --  free -h": "free -h && " + continueMsg,
	}
}
