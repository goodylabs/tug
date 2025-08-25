package docker

const dockerListCmd = "docker ps --format json"

var commandTemplates = map[string]string{
	"[docker]  logs <resource> -f":             "docker logs -f %s && read",
	"[docker]  exec -u root -ti <resource> sh": "docker exec -u root -it %s sh",
	"[docker]  logs <resource>":                "docker logs %s && read",
	"[docker]  restart <resource>":             "docker restart %s && read",
	"[docker]  ps":                             "watch docker ps %s",
	"[docker]  ps -a":                          "watch docker ps -a %s",
	"[bash]    bash":                           "bash",
	"[bash]    htop":                           "htop",
	"[django]  python manage.py shell":         "docker exec -u root -it %s python manage.py shell",
	"[traefik] show reverse-proxy config":      "docker exec %s sh -c 'apk add --no-cache --no-progress -q curl && curl localhost:8080/api/rawdata' | jq && read",
}
