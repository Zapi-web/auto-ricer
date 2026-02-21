# 🐧 Auto-Ricer: System Daemon for KDE/Linux

A lightweight system daemon that automatically updates your wallpaper and terminal color schemes (ricing) on the fly.

![License](https://img.shields.io/badge/License-MIT-blue.svg)

## 📦 Dependencies
This project requires:
* **Go** (for the daemon)
* **pipx** (to run pywal)
* **pywal** (installed via pipx):
```bash
pipx install pywal
```

## 🚀 Installation
    1. Clone the repository
```bash
    git clone https://github.com/Zapi-web/auto-ricer.git
    cd auto-ricer
```
    2. build the binary
```bash
    cd cmd/watchdog
    go build -o auto-ricer
```

## Usage
    You can run the daemon manually using flags:
    -dir - path to a dir with a pictures of your wallpaper
    -sh - path to a *.sh* script
    -lvl - level of logging (info, error, warn, debug; default: info)
    **Example:**
    ```bash
        ./auto-ricer -dir /home/USER/path/to/auto-ricer/pictures -sh /home/USER/path/to/auto-ricer/scripts/update_theme.sh -lvl error
        ```

## Enable Auto-run (systemd)
 1. Open systemd/auto-ricer.service.example and update the execution paths to match your system.
 2. Move and enable the service
 ```bash
    cd systemd
    mkdir -p ~/.config/systemd/user/
    mv auto-ricer.service.example ~/.config/systemd/user/auto-ricer.service
 ```
 start the auto-run
 ```bash
 systemctl --user daemon-reload
 systemctl --user enable auto-ricer
 systemctl --user start auto-ricer
 ```
 now check the status
 ```bash
 systemctl --user status auto-ricer
 ```
 and if it "Active: active (running)" you do all correct
