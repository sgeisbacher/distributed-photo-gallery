package main

import (
	"log"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
	"github.com/sgeisbacher/distributed-photo-gallery/helper"
	"github.com/sgeisbacher/distributed-photo-gallery/store"
)

func main() {
	cmdHandlerFactory := func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
		return []cqrs.CommandHandler{
			events.NoOpCommandHandler{eb},
		}
	}
	eventHandlerFactory := func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler {
		return []cqrs.EventHandler{
			store.CreateMediaOnMediaImportedHandler{cb},
		}
	}
	cqrsF := helper.CreateCqrsContext(cmdHandlerFactory, eventHandlerFactory)

	go func(cqrsCtx helper.CqrsContext) {
		if err := cqrsCtx.Run(); err != nil {
			panic(err)
		}
	}(cqrsF)

	log.Fatal(http.ListenAndServe("127.0.0.1:8788", nil))
}
