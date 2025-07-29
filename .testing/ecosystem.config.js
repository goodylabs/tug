module.exports = {
    apps: [
        {
            name: "pm2-app-1",
        },
    ],
    deploy: {
        staging: {
            user: "staging-user",
            host: "xxx.xxx.xxx.xxx",
        },
        staging_RO: {
            user: "staging-user",
            host: "yyy.yyy.yyy.yyy",
        },
        production_1: {
            user: "ubuntu",
            host: ["zzz.zzz.zzz.zzz"],
        },
        production_2: {
            user: "ubuntu",
            host: ["aaa.aaa.aaa.aaa", "ooo.ooo.ooo.ooo"],
        },
        production_RO_1: {
            user: "root",
            host: ["bbb.bbb.bbb.bbb"],
        },
        production_RO_2: {
            user: "root",
            host: ["ccc.ccc.ccc.ccc", "ddd.ddd.ddd.ddd", "eee.eee.eee.eee"],
        },
    },
};
