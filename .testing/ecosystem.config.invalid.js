module.exports = {
    apps: [
        {
            namez: "pm2-app-1",
            script: "",
        },
        {
            name: 123,
        },
    ],
    deploy: {
        staging: {
            user: "",
            host: "256.256.256.256",
        },
        staging_RO: {
            user: "staging-user",
        },
        production_1: {
            user: "ubuntu",
            host: "zzz.zzz.zzz.zzz",
            ref: "origin/main",
        },
        production_2: {
            host: ["aaa.aaa.aaa.aaa"],
        },
        production_RO_1: {
            user: "root",
            host: "bbb.bbb.bbb.bbb",
        },
        production_RO_2: {
            user: "root",
            host: ["ccc.ccc.ccc.ccc"],
            extra: true,
        },
        "": {
            user: "ghost",
            host: ["127.0.0.1"],
        },
    },
};
