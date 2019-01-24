package store

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
)

// todo var preStage map[string]Media
var db map[string]Media

func init() {
	db = make(map[string]Media)
}

func Print() {
	fmt.Println("db:")
	for _, media := range db {
		fmt.Printf("\t%#v\n", media)
	}
}

// CreateMediaOnMediaImportedHandler event handler
type CreateMediaOnMediaImportedHandler struct {
	CommandBus *cqrs.CommandBus
}

// NewEvent creates new event
func (h CreateMediaOnMediaImportedHandler) NewEvent() interface{} {
	return &events.MediaImported{}
}

// Handle handle event
func (h CreateMediaOnMediaImportedHandler) Handle(e interface{}) error {
	// todo race
	event := e.(*events.MediaImported)
	fmt.Printf("creating media %q in db\n", event.ID)
	if _, ok := db[event.ID]; ok {
		fmt.Println("WARN: already exists, overriding ...")
	}
	db[event.ID] = Media{
		ID:       event.ID,
		Name:     event.Name,
		OrigPath: event.Path,
		Type:     event.Type,
		Status:   created,
	}
	return nil
}
