package pm2

const (
	jlistCmd            = `source ~/.nvm/nvm.sh; pm2 jlist | sed -n '/^\[/,$p'`
	logsCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 logs %s`
	logsAllCmdTemplate  = `source ~/.nvm/nvm.sh; pm2 logs && %s read`
	showCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 show %s && read`
	restartCmdTemplate  = `source ~/.nvm/nvm.sh; pm2 restart %s`
	describeCmdTemplate = `source ~/.nvm/nvm.sh; pm2 describe %s && read`
	monitCmdTemplate    = `source ~/.nvm/nvm.sh; pm2 monit %s`
	updateCmdTemplate   = `source ~/.nvm/nvm.sh; pm2 update`
	bashCmdTemplate     = `source ~/.nvm/nvm.sh; bash && %s read`
	htopCmdTemplate     = `source ~/.nvm/nvm.sh; htop && %s read`
)

var commandTemplates = map[string]string{
	"[pm2] logs      <resource>": logsCmdTemplate,
	"[pm2] logs      <all>":      logsAllCmdTemplate,
	"[pm2] show      <resource>": showCmdTemplate,
	"[pm2] restart   <resource>": restartCmdTemplate,
	"[pm2] describe  <resource>": describeCmdTemplate,
	"[pm2] monit     <resource>": monitCmdTemplate,
	"[pm2] update":               updateCmdTemplate,
	"[shell] bash":               bashCmdTemplate,
	"[shell] htop":               htopCmdTemplate,
}
