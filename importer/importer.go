package importer

import (
	"os"
	"path/filepath"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
	"github.com/sirupsen/logrus"
)

func Run(rootDir string, cb *cqrs.CommandBus) error {
	logrus.Info("starting importer ...")
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.Errorf("error while walking %q: %v", path, err)
			logrus.Error("ignoring...")
			return nil
		}
		if info.IsDir() {
			return nil
		}
		logrus.Debug("found", path)
		cb.Send(events.ImportMedia{Path: path})
		return nil
	})
}