package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/helper"
	"github.com/sgeisbacher/distributed-photo-gallery/importer"
	"github.com/sgeisbacher/distributed-photo-gallery/stats"
	"github.com/sgeisbacher/distributed-photo-gallery/store"
)

var rootDir string

func init() {
	flag.StringVar(&rootDir, "p", "/tmp/photos", "path to scan for photos")
}

func main() {
	flag.Parse()

	// logrus.SetLevel(logrus.DebugLevel)

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
	go func(commandBus *cqrs.CommandBus) {
		time.Sleep(15 * time.Second)
		err := importer.Run(rootDir, commandBus)
		if err != nil {
			fmt.Printf("error while importing %q: %v\n", rootDir, err)
		}
	}(cqrsCtx.Facade.CommandBus())

	importer.Watch(cqrsCtx.Facade.CommandBus(), rootDir, false)

	go func(cqrsCtx helper.CqrsContext) {
		if err := cqrsCtx.Run(); err != nil {
			log.Fatal(err)
		}
	}(cqrsCtx)

	log.Fatal(http.ListenAndServe("127.0.0.1:8787", nil))
}
