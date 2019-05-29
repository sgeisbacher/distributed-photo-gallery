package events

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type NoOpCommandHandler struct {
	EventBus *cqrs.EventBus
}

func (h NoOpCommandHandler) NewCommand() interface{} {
	return &NoOp{}
}

func (h NoOpCommandHandler) Handle(ctx context.Context, c interface{}) error {
	return nil
}

func (h NoOpCommandHandler) HandlerName() string {
	return "NoOpCommandHandler"
}
