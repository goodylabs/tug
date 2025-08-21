package swarm

const swarmListCmd = "docker service ls --format json"

var commandTemplates = map[string]string{
	"[swarm] ps <service>":      "watch docker service ps %s --no-trunc",
	"[swarm] inspect <service>": "docker service inspect %s | jq && read",
	"[swarm] logs -f <service>": "docker service logs -f %s",
	"[bash]  bash":              "bash",
}
