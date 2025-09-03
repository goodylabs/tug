package docker

const (
	TARGET_IP_VAR    = "TARGET_IP"
	IP_ADDRESS_VAR   = "IP_ADDRESS"
	IP_ADDRESSES_VAR = "IP_ADDRESSES"
)

const dockerListCmd = "docker ps --format json"

const continueMsg = "echo 'Done, press Enter to continue...' && read"

var commandTemplates = map[string]string{
	"[docker]  logs -f          <resource>":        "docker logs -f %s",
	"[docker]  exec -u root -ti <resource>   sh":   "docker exec -u root -it %s sh",
	"[docker]  logs             <resource> | less": "docker logs %s | less",
	"[docker]  restart          <resource>":        "docker restart %s && " + continueMsg,
	"[docker]  stats":                              "docker stats",
	"[docker]  ps":                                 "watch docker ps",
	"[docker]  ps -a":                              "watch docker ps -a",
	"[bash]    bash":                               "bash",
	"[bash]    htop":                               "htop",
	"[django]  python manage.py shell":             "docker exec -u root -it %s python manage.py shell",
	"[traefik] show reverse-proxy config":          "docker exec %s sh -c 'apk add --no-cache --no-progress -q curl && curl localhost:8080/api/rawdata' | jq && " + continueMsg,
}
