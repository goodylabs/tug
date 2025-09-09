package services

const continueMsg = "echo 'Done, press Enter to continue...' && read"

func GetActionTemplates() map[string]string {
	return map[string]string{
		"[swarm] ls":                                 "watch docker service ls",
		"[swarm] ps      <service>":                  "watch docker service ps %s --no-trunc",
		"[swarm] ps      <service>   (only running)": `watch 'docker service ps --filter desired-state=running --format "{{.ID}} {{.Name}} - {{.Node}} | {{.Image}}" %s'`,
		"[swarm] inspect <service>":                  "docker service inspect %s | jq && read",
		"[swarm] update  <service>":                  "docker service update %s --force && " + continueMsg,
		"[swarm] logs -f <service>":                  "docker service logs -f %s",
		"[swarm] logs    <service> | less":           "docker service logs %s | less",
		"[swarm] scale   <service>   replicas to 0":  "docker service scale %s=0 && " + continueMsg,
		"[swarm] scale   <service>   replicas to 1":  "docker service scale %s=1 && " + continueMsg,
		"[swarm] scale   <service>   replicas to 3":  "docker service scale %s=3 && " + continueMsg,
		"[bash]  bash":                               "bash",
		// "[swarm] remove  <service>":                  "docker service remove %s && " + continueMsg,
	}
}
