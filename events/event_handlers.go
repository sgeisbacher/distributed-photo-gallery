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

type MediaThumbShotGeneratedHandler struct {
	commandBus *cqrs.CommandBus
}

func (h MediaThumbShotGeneratedHandler) NewEvent() interface{} {
	return &MediaThumbShotGenerated{}
}

func (h MediaThumbShotGeneratedHandler) Handle(e interface{}) error {
	fmt.Println("MediaThumbShotGeneratedHandler eventhandler")
	return nil
}

/*
*
*      MediaBigShotGenerated-Eventhandler
*
 */

type MediaBigShotGeneratedHandler struct {
	commandBus *cqrs.CommandBus
}

func (h MediaBigShotGeneratedHandler) NewEvent() interface{} {
	return &MediaBigShotGenerated{}
}

func (h MediaBigShotGeneratedHandler) Handle(e interface{}) error {
	fmt.Println("MediaBigShotGenerated eventhandler")
	return nil
}
