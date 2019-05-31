package stats

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
	"github.com/sirupsen/logrus"
)

type TrackStatsOnMediaImportedHandler struct {
	CommandBus *cqrs.CommandBus
	Store      *StatsStore
}

func (h TrackStatsOnMediaImportedHandler) NewEvent() interface{} {
	return &events.MediaImported{}
}

func (h TrackStatsOnMediaImportedHandler) Handle(ctx context.Context, e interface{}) error {
	event := e.(*events.MediaImported)
	logrus.Debug("tracking stats for media:", event.Path)
	h.Store.Track(event.Type, event.Size)
	return nil
}

func (h TrackStatsOnMediaImportedHandler) HandlerName() string {
	return "TrackStatsOnMediaImportedHandler"
}
