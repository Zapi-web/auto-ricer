package main

import (
	"log"

	"github.com/Zapi-web/auto-ricer/internal/watcher"
)

func main() {
	path := "~/Documents/auto-ricer/pictures"
	
	watch, err := watcher.NewWatcher(path)

	if err != nil {
		log.Fatal("Failed to create a watcher: %w", err)
	}
	defer watch.Close()

	err = watch.Watch()

	if err != nil {
		log.Fatal("Failed to watch a dir: %w", err)
	}
}
