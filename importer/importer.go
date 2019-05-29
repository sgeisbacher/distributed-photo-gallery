package importer

import (
	"context"
	"os"
	"path/filepath"
	"time"

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
		cb.Send(context.Background(), events.ImportMedia{Path: path})
		time.Sleep(3 * time.Second)
		return nil
	})
}
