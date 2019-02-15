# Registrybot

```
[Unit]
Description=login.service
Requires=docker.service
After=docker.service
[Timer]
OnCalendar=00/3:00:00
[Service]
TimeoutStartSec=0
Type=oneshot
RemainAfterExit=yes
ExecStartPre=-/usr/bin/docker kill registrybot
ExecStartPre=-/usr/bin/docker rm registrybot
ExecStartPre=-/usr/bin/docker pull jayhanjaelee/registrybot:latest
ExecStart=-/bin/bash -c 'eval $(/usr/bin/docker run -e AWS_ACCESS_KEY_ID="" -e AWS_SECRET_ACCESS_KEY="" -e AWS_DEFAULT_REGION="" -e AWS_ECS_REGISTRY_ID="" --rm --name registrybot jayhanjaelee/registrybot:latest)'
[Install]
WantedBy=timers.target
```
