package services

import "github.com/goodylabs/tug/internal/modules"

const continueMsg = "echo 'Done, press Enter to continue...' && read"

func GetActionTemplates() []modules.TechCmdTemplate {
	return []modules.TechCmdTemplate{
		{
			Display:  "[docker]  logs -f          <resource>",
			Template: "docker logs -f %s",
		},
		{
			Display:  "[docker]  exec -u root -ti <resource>   sh",
			Template: "docker exec -u root -it %s sh",
		},
		{
			Display:  "[docker]  logs             <resource> | less",
			Template: "docker logs %s | less",
		},
		{
			Display:  "[docker]  restart          <resource>",
			Template: "docker restart %s && " + continueMsg,
		},
		{
			Display:  "[docker]  stats",
			Template: "docker stats",
		},
		{
			Display:  "[docker]  ps",
			Template: "watch docker ps",
		},
		{
			Display:  "[docker]  ps -a",
			Template: "watch docker ps -a",
		},
		{
			Display:  "[bash]    bash",
			Template: "bash",
		},
		{
			Display:  "[bash]    htop",
			Template: "htop",
		},
		{
			Display:  "[django]  python manage.py shell",
			Template: "docker exec -u root -it %s python manage.py shell",
		},
		{
			Display:  "[traefik] show reverse-proxy config",
			Template: "docker exec %s sh -c 'apk add --no-cache --no-progress -q curl && curl localhost:8080/api/rawdata' | jq && " + continueMsg,
		},
	}
}
