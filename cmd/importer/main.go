package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/helper"
	"github.com/sgeisbacher/distributed-photo-gallery/importer"
)

var path string

func init() {
	flag.StringVar(&path, "p", "/tmp/photos", "path to scan for photos")
}

func main() {
	flag.Parse()

	cmdHandlerFactory := func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
		return []cqrs.CommandHandler{
			importer.ImportMediaHandler{eb},
		}
	}
	eventHandlerFactory := func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler {
		return []cqrs.EventHandler{}
	}
	cqrsCtx := helper.CreateCqrsContext(cmdHandlerFactory, eventHandlerFactory)

	// starting importer
	go func(commandBus *cqrs.CommandBus) {
		time.Sleep(15 * time.Second)
		err := importer.Run(path, commandBus)
		if err != nil {
			fmt.Printf("error while importing %q: %v\n", path, err)
		}
	}(cqrsCtx.Facade.CommandBus())

	importer.Watch(cqrsCtx.Facade.CommandBus(), path, false)

	if err := cqrsCtx.Run(); err != nil {
		panic(err)
	}
}
