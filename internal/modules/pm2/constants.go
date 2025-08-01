package pm2

const (
	jlistCmd            = `source ~/.nvm/nvm.sh; pm2 jlist | sed -n '/^\[/,$p'`
	logsCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 logs %s`
	showCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 show %s && read`
	restartCmdTemplate  = `source ~/.nvm/nvm.sh; pm2 restart %s`
	describeCmdTemplate = `source ~/.nvm/nvm.sh; pm2 describe %s && read`
	monitCmdTemplate    = `source ~/.nvm/nvm.sh; pm2 monit %s`
	updateCmdTemplate   = `source ~/.nvm/nvm.sh; pm2 update`
)

var commandTemplates = map[string]string{
	"[pm2] logs <resource>":     logsCmdTemplate,
	"[pm2] show <resource>":     showCmdTemplate,
	"[pm2] restart <resource>":  restartCmdTemplate,
	"[pm2] describe <resource>": describeCmdTemplate,
	"[pm2] monit <resource>":    monitCmdTemplate,
	"[pm2] update":              updateCmdTemplate,
}

const tmpJsonPath = "/tmp/ecosystem.json"

const ecosystemJsScript = `const fs = require("fs");
const config = require("%s");

if (config.deploy) {
	for (const key in config.deploy) {
		const deployEntry = config.deploy[key];
		if (typeof deployEntry.host === "string") {
			deployEntry.host = [deployEntry.host];
		}
	}
}

fs.writeFileSync("%s", JSON.stringify(config, null, 2));`

const ecosystemCjsScript = `import config from "%s";

const { default: ecosystemConfig } = await import("%s");

if (ecosystemConfig.deploy) {
	for (const key in ecosystemConfig.deploy) {
		const deployEntry = ecosystemConfig.deploy[key];
		if (typeof deployEntry.host === "string") {
			deployEntry.host = [deployEntry.host];
		}
	}
}

const fs = await import("fs/promises");
await fs.writeFile("%s", JSON.stringify(ecosystemConfig, null, 2));`
