package services

const continueMsg = "echo 'Done, press Enter to continue...' && read"

func GetActionTemplates() map[string]string {
	return map[string]string{
		"docker   --  logs -f          <resource>":      "docker logs -f %s",
		"docker   --  exec -u root -ti <resource>   sh": "docker exec -u root -it %s sh",
		// "docker   --  logs             <resource> | less": "docker logs %s | less",
		// "docker   --  restart          <resource>":        "docker restart %s && " + continueMsg,
		// "docker   --  stats":                              "docker stats",
		// "docker   --  ps":                                 "watch docker ps",
		// "docker   --  ps -a":                              "watch docker ps -a",
		// "docker   --  inspect          <resource>":        "docker inspect %s | jq | less",
		// "bash     --  bash":                               "bash",
		// "bash     --  htop":                               "htop",
		// "django   --  python manage.py shell":             "docker exec -u root -it %s python manage.py shell",
		// "django   --  python manage.py shell_plus":        "docker exec -u root -it %s python manage.py shell_plus",
		// "traefik  --  show config routers":                "docker exec %s sh -c 'apk add --no-cache --no-progress -q curl && curl localhost:8080/api/rawdata' | jq '.routers' | less",
		// "traefik  --  show config services":               "docker exec %s sh -c 'apk add --no-cache --no-progress -q curl && curl localhost:8080/api/rawdata' | jq '.services' | less",
	}
}
