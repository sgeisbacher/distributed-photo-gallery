package main

import (
	"log"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/helper"
	"github.com/sgeisbacher/distributed-photo-gallery/stats"
)

func main() {
	statsStore := stats.NewStatsStore()

	cmdHandlerFactory := func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
		return []cqrs.CommandHandler{}
	}
	eventHandlerFactory := func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler {
		return []cqrs.EventHandler{
			stats.TrackStatsOnMediaImportedHandler{cb, statsStore},
		}
	}
	cqrsF := helper.CreateCqrsContext(cmdHandlerFactory, eventHandlerFactory)

	go func(cqrsCtx helper.CqrsContext) {
		if err := cqrsCtx.Run(); err != nil {
			panic(err)
		}
	}(cqrsF)

	log.Fatal(http.ListenAndServe("127.0.0.1:8787", nil))
}
