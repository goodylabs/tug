package services

import "github.com/goodylabs/tug/internal/modules"

const continueMsg = "echo 'Done, press Enter to continue...' && read"

func GetActionTemplates() []modules.TechCmdTemplate {
	return []modules.TechCmdTemplate{
		{
			Display:  "[swarm] ls",
			Template: "watch docker service ls",
		},
		{
			Display:  "[swarm] ps      <service>",
			Template: "watch docker service ps %s --no-trunc",
		},
		{
			Display:  "[swarm] ps      <service>   (only running)",
			Template: `watch 'docker service ps --filter desired-state=running --format "{{.ID}} {{.Name}} - {{.Node}} | {{.Image}}" %s'`},
		{
			Display:  "[swarm] inspect <service>",
			Template: "docker service inspect %s | jq && read"},
		{
			Display:  "[swarm] update  <service>",
			Template: "docker service update %s --force && " + continueMsg,
		},
		{
			Display:  "[swarm] logs -f <service>",
			Template: "docker service logs -f %s",
		},
		{
			Display:  "[swarm] logs    <service> | less",
			Template: "docker service logs %s | less",
		},
		{
			Display:  "[swarm] scale   <service>   replicas to 0",
			Template: "docker service scale %s=0 && " + continueMsg,
		},
		{
			Display:  "[swarm] scale   <service>   replicas to 1",
			Template: "docker service scale %s=1 && " + continueMsg,
		},
		{
			Display:  "[swarm] scale   <service>   replicas to 3",
			Template: "docker service scale %s=3 && " + continueMsg,
		},
		{
			Display:  "[bash]  bash",
			Template: "bash",
		},
	}
}
