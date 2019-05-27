package events

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type NoOpCommandHandler struct {
	EventBus *cqrs.EventBus
}

func (h NoOpCommandHandler) NewCommand() interface{} {
	return &NoOp{}
}

func (h NoOpCommandHandler) Handle(c interface{}) error {
	return nil
}
