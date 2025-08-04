package pm2

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
