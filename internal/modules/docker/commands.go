package docker

const (
	dockerListCmd     = "docker ps --format json"
	dockerLogsLiveCmd = "docker logs -f %s && read"
	dockerExecCmd     = "docker exec -it %s sh"
	dockerLogsCmd     = "docker logs %s && read"
	dockerRestartCmd  = "docker restart %s && read"
	bashCmd           = "bash"
	htopCmd           = "htop"
)

var commandTemplates = map[string]string{
	"[docker] logs <resource> -f":     dockerLogsLiveCmd,
	"[docker] exec -ti <resource> sh": dockerExecCmd,
	"[docker] logs <resource>":        dockerLogsCmd,
	"[docker] restart <resource>":     dockerRestartCmd,
	"[bash]   bash":                   bashCmd,
	"[bash]   htop":                   htopCmd,
}
