package pm2_test

const output1 = `[{ "name": "pm2-logrotate", "pid": 3105},{  "name": "staging_ro", "pid": 4043331},{ "name": "api-staging", "pid": 4041884}]`

const invalidOutput1 = `>>>> In-memory PM2 is out-of-date, do:
>>>> $ pm2 update
In memory PM2 version: 5.2.2
Local PM2 version: 6.0.8

[{ "name": "pm2-logrotate", "pid": 3105},{  "name": "staging_ro", "pid": 4043331},{ "name": "api-staging", "pid": 4041884}]`
