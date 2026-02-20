package main

import (
	"flag"
	"os"
	"strings"

	"github.com/Zapi-web/auto-ricer/internal/executor"
	"github.com/Zapi-web/auto-ricer/internal/logger"
	"github.com/Zapi-web/auto-ricer/internal/watcher"
)

func main() {
	path := flag.String("dir", "./pictures", "Path to wallpaper dir")
	pathToSh := flag.String("sh", "./scripts/update_theme.sh", "Path to a script")	
	level := flag.String("lvl", "info", "logger level")

	flag.Parse()
	
	lowLvl := strings.ToLower(*level)
	logger.NewLogger(lowLvl)
	logger.Log.Info("Logger initialized", "level", lowLvl)
	
	watch, err := watcher.NewWatcher(*path)


	if err != nil {
		logger.Log.Error("Failed to create a watcher", "error", err)
		os.Exit(1)
	}
	defer watch.Close()
	logger.Log.Info("Sucsesfully created a watcher")

	err = watch.Watch()

	if err != nil {
		logger.Log.Error("Failed to watch a dir", "dir", *path, "error", err)
		os.Exit(1)
	}

	executor.UpdateTheme(watch.Events, *pathToSh)
}
