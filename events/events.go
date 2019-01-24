package events

import "github.com/ThreeDotsLabs/watermill/components/cqrs"

// MediaImported event
type MediaImported struct {
	commandBus *cqrs.CommandBus
	ID         string
	Path       string
	Name       string
	Size       int64
	Type       string
}

// MediaThumbShotGenerated event
type MediaThumbShotGenerated struct {
	commandBus *cqrs.CommandBus
	ID         string
	Path       string
}

// MediaBigShotGenerated event
type MediaBigShotGenerated struct {
	commandBus *cqrs.CommandBus
	ID         string
	Path       string
}

// MediaApproved event
type MediaApproved struct {
	commandBus *cqrs.CommandBus
	ID         string
}
