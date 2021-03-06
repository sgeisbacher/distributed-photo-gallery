package convertion

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
)

// GenerateThumbAndBigShotsOnMediaImportedHandler event handler
type GenerateThumbAndBigShotsOnMediaImportedHandler struct {
	CommandBus *cqrs.CommandBus
}

// NewEvent creates new event
func (h GenerateThumbAndBigShotsOnMediaImportedHandler) NewEvent() interface{} {
	return &events.MediaImported{}
}

// Handle handle event
func (h GenerateThumbAndBigShotsOnMediaImportedHandler) Handle(e interface{}) error {
	event := e.(*events.MediaImported)
	fmt.Printf("media %q has been imported, triggering shots-generation...\n", event.Path)
	// h.CommandBus.Send(GenerateThumbShot{})
	// h.CommandBus.Send(GenerateBigShot{})
	return nil
}
