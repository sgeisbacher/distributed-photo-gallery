package helper

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/infrastructure/gochannel"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

type CqrsContext struct {
	Facade *cqrs.Facade
	Run    func() error
}

type CommandHandlerFactory func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler
type EventHandlerFactory func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler

func CreateCqrsContext(cmdHandlerFactory CommandHandlerFactory, eventHandlerFactory EventHandlerFactory) CqrsContext {
	logger := watermill.NewStdLogger(false, false)
	marshaler := cqrs.JSONMarshaler{}

	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddMiddleware(middleware.Recoverer)

	cqrsFacade, err := cqrs.NewFacade(cqrs.FacadeConfig{
		CommandsTopic:         "commands",
		EventsTopic:           "events",
		CommandHandlers:       cmdHandlerFactory,
		EventHandlers:         eventHandlerFactory,
		Router:                router,
		CommandsPubSub:        pubSub,
		EventsPubSub:          pubSub,
		Logger:                logger,
		CommandEventMarshaler: marshaler,
	})

	return CqrsContext{
		Facade: cqrsFacade,
		Run: func() error {
			return router.Run()
		},
	}
}
