package services

import "github.com/goodylabs/tug/internal/modules"

const continueMsg = "echo 'Done, press Enter to continue...' && read"

func GetActionTemplates() []modules.TechCmdTemplate {
	return []modules.TechCmdTemplate{
		{
			Display:  "[pm2]  logs     <resource> | less",
			Template: `source ~/.nvm/nvm.sh; pm2 logs %s | less`,
		},
		{
			Display:  "[pm2]  logs     <resource>",
			Template: `source ~/.nvm/nvm.sh; pm2 logs %s`,
		},
		{
			Display:  "[pm2]  logs",
			Template: `source ~/.nvm/nvm.sh; pm2 logs`,
		},
		{
			Display:  "[pm2]  show     <resource>",
			Template: `source ~/.nvm/nvm.sh; pm2 show %s && read`,
		},
		{
			Display:  "[pm2]  restart  <resource>",
			Template: `source ~/.nvm/nvm.sh; pm2 restart %s`,
		},
		{
			Display:  "[pm2]  describe <resource>",
			Template: `source ~/.nvm/nvm.sh; pm2 describe %s && read`,
		},
		{
			Display:  "[pm2]  monit",
			Template: `source ~/.nvm/nvm.sh; pm2 monit`,
		},
		{
			Display:  "[pm2]  update",
			Template: `source ~/.nvm/nvm.sh; pm2 update`,
		},
		{
			Display:  "[bash] bash",
			Template: `source ~/.nvm/nvm.sh; bash`,
		},
		{
			Display:  "[bash] htop",
			Template: `source ~/.nvm/nvm.sh; htop`,
		},
	}
}
