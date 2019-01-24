package importer

import (
	"fmt"
	"path"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
)

// ImportMediaHandler command handler
type ImportMediaHandler struct {
	EventBus *cqrs.EventBus
}

// NewCommand creates new ImportMedia command
func (h ImportMediaHandler) NewCommand() interface{} {
	return &events.ImportMedia{}
}

// Handle handles ImportMedia command
func (h ImportMediaHandler) Handle(c interface{}) error {
	cmd, ok := c.(*events.ImportMedia)
	if !ok {
		return fmt.Errorf("could not typeassert")
	}
	fmt.Printf("importing %q ...\n", cmd.Path)
	id := uuid.New().String()
	name := path.Base(cmd.Path)
	fileType := path.Ext(cmd.Path)
	h.EventBus.Publish(events.MediaImported{
		ID:   id,
		Name: name,
		Path: cmd.Path,
		Type: fileType[1:],
	})
	return nil
}
