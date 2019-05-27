package importer

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
	"github.com/sirupsen/logrus"
)

// ImportMediaHandler command handler
type ImportMediaHandler struct {
	EventBus *cqrs.EventBus
}

// NewCommand creates new ImportMedia command
func (h ImportMediaHandler) NewCommand() interface{} {
	return &events.ImportMedia{}
}

func isValid(path string, info os.FileInfo) (bool, string) {
	lpath := strings.ToLower(path)
	if strings.HasSuffix(lpath, ".ds_store") {
		return false, fmt.Sprintf("ignoring thumb-file: %s", path)
	}
	if strings.HasSuffix(lpath, "thumbs.db") {
		return false, fmt.Sprintf("ignoring thumb-file: %s", path)
	}
	if info.IsDir() {
		return false, fmt.Sprintf("ignoring directory: %s", path)
	}
	return true, ""
}

// Handle handles ImportMedia command
func (h ImportMediaHandler) Handle(c interface{}) error {
	cmd, ok := c.(*events.ImportMedia)
	if !ok {
		return fmt.Errorf("could not typeassert")
	}
	info, err := os.Stat(cmd.Path)
	if err != nil {
		logrus.Errorf("ERROR: could not stat file %q, ignoring: %v", cmd.Path, err)
		return nil
	}
	if ok, reason := isValid(cmd.Path, info); !ok {
		logrus.Warn(reason)
		return nil
	}
	id := uuid.New().String()
	logger := logrus.WithField("media-id", id)
	logger.Infof("importing %q ...", cmd.Path)
	name := path.Base(cmd.Path)
	fileType := path.Ext(cmd.Path)
	logger.Debugf("file: %s, fileType: %s", cmd.Path, fileType)
	h.EventBus.Publish(events.MediaImported{
		ID:   id,
		Name: name,
		Path: cmd.Path,
		Type: fileType[1:],
		Size: info.Size(),
	})
	return nil
}
