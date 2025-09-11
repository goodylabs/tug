package pm2

const (
	jlistCmd = `source ~/.nvm/nvm.sh; pm2 jlist | sed -n '/^\[/,$p'`
)

var commandTemplates = map[string]string{
	"[pm2]  logs     <resource> | less": `source ~/.nvm/nvm.sh; pm2 logs %s | less`,
	"[pm2]  logs     <resource>":        `source ~/.nvm/nvm.sh; pm2 logs %s`,
	"[pm2]  logs":                       `source ~/.nvm/nvm.sh; pm2 logs`,
	"[pm2]  show     <resource>":        `source ~/.nvm/nvm.sh; pm2 show %s && read`,
	"[pm2]  restart  <resource>":        `source ~/.nvm/nvm.sh; pm2 restart %s`,
	"[pm2]  describe <resource>":        `source ~/.nvm/nvm.sh; pm2 describe %s && read`,
	"[pm2]  monit":                      `source ~/.nvm/nvm.sh; pm2 monit`,
	"[pm2]  update":                     `source ~/.nvm/nvm.sh; pm2 update`,
	"[bash] bash":                       `source ~/.nvm/nvm.sh; bash`,
	"[bash] htop":                       `source ~/.nvm/nvm.sh; htop`,
}
