package main

import (
	"time"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/helper"
	"github.com/sgeisbacher/distributed-photo-gallery/stats"
)

func main() {
	cmdHandlerFactory := func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
		return []cqrs.CommandHandler{}
	}
	eventHandlerFactory := func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler {
		return []cqrs.EventHandler{
			stats.TrackStatsOnMediaImportedHandler{cb},
		}
	}
	cqrsF := helper.CreateCqrsContext(cmdHandlerFactory, eventHandlerFactory)

	time.AfterFunc(10*time.Second, stats.Print)

	if err := cqrsF.Run(); err != nil {
		panic(err)
	}
}
