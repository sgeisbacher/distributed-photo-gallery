package importer

import (
	"log"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
	"gopkg.in/fsnotify.v1"
)

func Watch(cb *cqrs.CommandBus, rootDir string, recursive bool) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go func(commandBus *cqrs.CommandBus) {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("new file:", event.Name)
					commandBus.Send(events.ImportMedia{Path: event.Name})
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}(cb)

	return watcher.Add(rootDir)
}
