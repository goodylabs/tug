package swarm

const swarmListCmd = "docker service ls --format json"

var commandTemplates = map[string]string{
	"[swarm] ls":                     "watch docker service ls",
	"[swarm] ps <service>":           "watch docker service ps %s --no-trunc",
	"[swarm] inspect <service>":      "docker service inspect %s | jq && read",
	"[swarm] remove <service>":       "docker service remove %s",
	"[swarm] force update <service>": "docker service update %s --force",
	"[swarm] logs -f <service>":      "docker service logs -f %s",
	"[bash]  bash":                   "bash",
}
