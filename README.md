# Clashadm: A CLI-Based Clash Management tool

## Features

- Install and Uninstall Clash in Linux
- Auto generate Systemd unit file
- Auto subscribe and update clash config file, and modify config and rules
- Config system to support Tun, redir for Gateway Proxy
- ~~Buildin Dashboard~~

## Status

This project is just getting started. and features are being implemented.

- [x] Install Clash
- [x] Config Subscription
- [x] System configuration
- [ ] Dashboard

And maybe TO-DO features

- [ ] GUI
- [ ] Non-Systemd system support
- [ ] ~~Restful-APIs~~

## Getting Started

Now you need to clone the source code and compile by yourself.

```shell
# Install Go first
make linux-amd64
./bin/clashadm-linux-amd64
```

## Thanks to

- Dreamacro/clash
- spf13/cobra
- coreos/go-systemd
