package pm2

const (
	jlistCmd            = `source ~/.nvm/nvm.sh; pm2 jlist | sed -n '/^\[/,$p'`
	logsCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 logs %s | less`
	logsAllCmdTemplate  = `source ~/.nvm/nvm.sh; pm2 logs`
	showCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 show %s && read`
	restartCmdTemplate  = `source ~/.nvm/nvm.sh; pm2 restart %s`
	describeCmdTemplate = `source ~/.nvm/nvm.sh; pm2 describe %s && read`
	monitCmdTemplate    = `source ~/.nvm/nvm.sh; pm2 monit`
	updateCmdTemplate   = `source ~/.nvm/nvm.sh; pm2 update`
	bashCmdTemplate     = `source ~/.nvm/nvm.sh; bash`
	htopCmdTemplate     = `source ~/.nvm/nvm.sh; htop`
)

var commandTemplates = map[string]string{
	"[pm2]  logs     <resource> | less": logsCmdTemplate,
	"[pm2]  logs":                       logsAllCmdTemplate,
	"[pm2]  show     <resource>":        showCmdTemplate,
	"[pm2]  restart  <resource>":        restartCmdTemplate,
	"[pm2]  describe <resource>":        describeCmdTemplate,
	"[pm2]  monit":                      monitCmdTemplate,
	"[pm2]  update":                     updateCmdTemplate,
	"[bash] bash":                       bashCmdTemplate,
	"[bash] htop":                       htopCmdTemplate,
}
