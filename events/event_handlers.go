package events

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

/*
*
*      MediaThumbShotGenerated-Eventhandler
*
 */

// MediaThumbShotGeneratedHandler event handler
type MediaThumbShotGeneratedHandler struct {
	commandBus *cqrs.CommandBus
}

// NewEvent creates new event
func (h MediaThumbShotGeneratedHandler) NewEvent() interface{} {
	return &MediaThumbShotGenerated{}
}

// Handle handle event
func (h MediaThumbShotGeneratedHandler) Handle(e interface{}) error {
	fmt.Println("MediaThumbShotGeneratedHandler eventhandler")
	return nil
}

/*
*
*      MediaBigShotGenerated-Eventhandler
*
 */

// MediaBigShotGeneratedHandler event handler
type MediaBigShotGeneratedHandler struct {
	commandBus *cqrs.CommandBus
}

// NewEvent creates new event
func (h MediaBigShotGeneratedHandler) NewEvent() interface{} {
	return &MediaBigShotGenerated{}
}

// Handle handle event
func (h MediaBigShotGeneratedHandler) Handle(e interface{}) error {
	fmt.Println("MediaBigShotGenerated eventhandler")
	return nil
}
