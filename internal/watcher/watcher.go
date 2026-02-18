package watcher

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
)

type Watcher struct {
	watcher *fsnotify.Watcher
	events chan string
	path string
}

func NewWatcher(path string) (*Watcher,error) {
	watcher, err := fsnotify.NewWatcher()
	
	if err != nil {
		return nil, fmt.Errorf("Error to create watcher: %w", err)
	}

	expandedPath, err := homedir.Expand(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to expand path: %w", err)
	}

	return &Watcher{
		watcher: watcher,
		events: make(chan string),
		path: expandedPath,
	}, nil
}

func (w *Watcher) Watch() error {
	go func() {
	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}
			if event.Has(fsnotify.Write) {
				log.Println(event.Name)
			}
		case err, ok := <-w.watcher.Errors:
			if !ok {
				return
			}
			log.Println(err)
		}
	}
}()
	err := w.watcher.Add(w.path)

	if err != nil {
		return fmt.Errorf("failed to add a watcher")
	}

	return nil
}

func (w *Watcher) Close() {
	w.watcher.Close()
}


