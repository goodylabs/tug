package services

import "github.com/goodylabs/tug/internal/modules"

const continueMsg = "echo 'Done, press Enter to continue...' && read"

var baseCmds = map[string]string{
	"docker  --  logs -f          <resource>":        "docker logs -f %s",
	"docker  --  exec -u root -ti <resource>   sh":   "docker exec -u root -it %s sh",
	"docker  --  logs             <resource> | less": "docker logs %s | less",
	"docker  --  restart          <resource>":        "docker restart %s && " + continueMsg,
	"docker  --  stats":                              "docker stats",
	"docker  --  ps":                                 "watch docker ps",
	"docker  --  ps -a":                              "watch docker ps -a",
	"bash    --  bash":                               "bash",
	"bash    --  htop":                               "htop",
	"django  --  python manage.py shell":             "docker exec -u root -it %s python manage.py shell",
}

var specificCmds = []modules.TechCmdTemplate{
	{
		Display:  "traefik --  show reverse-proxy config",
		Template: "docker exec %s sh -c 'apk add --no-cache --no-progress -q curl && curl localhost:8080/api/rawdata' | jq && " + continueMsg,
		Filter:   "traefik* --",
	},
}

func GetActionTemplates() []modules.TechCmdTemplate {
	var result []modules.TechCmdTemplate
	for key, value := range baseCmds {
		result = append(result, modules.TechCmdTemplate{
			Display:  key,
			Template: value,
		})
	}
	result = append(result, specificCmds...)
	return result
}
