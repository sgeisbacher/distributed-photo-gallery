package helper

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/infrastructure/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

const projectID = "theta-disk-241906"

type CqrsContext struct {
	Facade *cqrs.Facade
	Run    func() error
}

type CommandHandlerFactory func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler
type EventHandlerFactory func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler

func CreateCqrsContext(cmdHandlerFactory CommandHandlerFactory, eventHandlerFactory EventHandlerFactory) CqrsContext {
	logger := watermill.NewStdLogger(false, false)
	marshaler := cqrs.JSONMarshaler{}

	googleCloudPub, err := googlecloud.NewPublisher(
		context.Background(),
		googlecloud.PublisherConfig{
			ProjectID: projectID,
			Logger:    watermill.NewStdLogger(false, false),
		},
	)
	if err != nil {
		panic(err)
	}

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddMiddleware(middleware.Recoverer)

	cqrsFacade, err := cqrs.NewFacade(cqrs.FacadeConfig{
		GenerateCommandsTopic: func(commandName string) string { return "commands" },
		GenerateEventsTopic:   func(eventName string) string { return "events" },
		CommandHandlers:       cmdHandlerFactory,
		EventHandlers:         eventHandlerFactory,
		EventsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			return createSubscriber(handlerName)
		},
		Router:            router,
		CommandsPublisher: googleCloudPub,
		EventsPublisher:   googleCloudPub,
		CommandsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			return createSubscriber(handlerName)
		},
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

func createSubscriber(handlerName string) (message.Subscriber, error) {
	return googlecloud.NewSubscriber(
		context.Background(),
		googlecloud.SubscriberConfig{
			GenerateSubscriptionName: func(topic string) string {
				return fmt.Sprintf("%s-%s", topic, handlerName)
			},
			ProjectID: projectID,
		},
		watermill.NewStdLogger(false, false),
	)
}
