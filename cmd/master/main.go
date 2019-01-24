package main

import (
	"flag"
	"time"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
	"github.com/sgeisbacher/distributed-photo-gallery/helper"
	"github.com/sgeisbacher/distributed-photo-gallery/importer"
	"github.com/sgeisbacher/distributed-photo-gallery/stats"
	"github.com/sgeisbacher/distributed-photo-gallery/store"
)

var path string

func init() {
	flag.StringVar(&path, "p", "", "path to scan for photos")
}

func main() {
	cmdHandlerFactory := func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
		return []cqrs.CommandHandler{
			importer.ImportMediaHandler{eb},
		}
	}
	eventHandlerFactory := func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler {
		return []cqrs.EventHandler{
			store.CreateMediaOnMediaImportedHandler{cb},
			stats.TrackStatsOnMediaImportedHandler{cb},
		}
	}
	cqrsCtx := helper.CreateCqrsContext(cmdHandlerFactory, eventHandlerFactory)

	// starting importer
	go func(ctx helper.CqrsContext) {
		filePaths := []string{"/tmp/photo1.jpg", "/tmp/photo2.png", "/tmp/photo3.png"}
		for _, fPath := range filePaths {
			ctx.Facade.CommandBus().Send(events.ImportMedia{Path: fPath})
			time.Sleep(1 * time.Second)
		}
	}(cqrsCtx)

	time.AfterFunc(10*time.Second, stats.Print)
	time.AfterFunc(15*time.Second, store.Print)

	if err := cqrsCtx.Run(); err != nil {
		panic(err)
	}
}
