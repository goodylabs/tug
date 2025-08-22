package swarm

const swarmListCmd = "docker service ls --format json"

var commandTemplates = map[string]string{
	"[swarm] ls": "watch docker service ls",

	"[swarm] ps <service>":      "watch docker service ps %s --no-trunc",
	"[swarm] inspect <service>": "docker service inspect %s | jq && read",

	"[swarm] remove <service>":       "docker service remove %s",
	"[swarm] force update <service>": "docker service update %s --force",
	"[swarm] logs -f <service>":      "docker service logs -f %s",

	"[swarm] scale <service> replicas to 0": "docker scale %s=0 && read",
	"[swarm] scale <service> replicas to 1": "docker scale %s=1 && read",
	"[swarm] scale <service> replicas to 3": "docker scale %s=3 && read",

	"[bash]  bash": "bash",
}
