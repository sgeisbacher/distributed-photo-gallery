package stats

import (
	"context"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
	"github.com/sgeisbacher/distributed-photo-gallery/helper"
	"github.com/sirupsen/logrus"
)

type statistics struct {
	counter          int64
	totalSize        int64
	typeDistribution map[string]int
}

type StatsDto struct {
	Counter          int64
	TotalSize        int64
	AvgSize          int64
	TypeDistribution map[string]int
}

type operation func(*statistics)

type TrackStatsOnMediaImportedHandler struct {
	CommandBus *cqrs.CommandBus
	ops        chan operation
}

func NewTrackStatsOnMediaImportedHandler(cb *cqrs.CommandBus) TrackStatsOnMediaImportedHandler {
	handler := TrackStatsOnMediaImportedHandler{
		CommandBus: cb,
		ops:        make(chan operation),
	}
	http.HandleFunc("/stats", handler.HandleGetStats)
	go handler.run()
	return handler
}

func (h TrackStatsOnMediaImportedHandler) NewEvent() interface{} {
	return &events.MediaImported{}
}

func (h TrackStatsOnMediaImportedHandler) Handle(ctx context.Context, e interface{}) error {
	event := e.(*events.MediaImported)
	logrus.Debug("tracking stats for media:", event.Path)

	h.ops <- func(st *statistics) {
		incTypeDistribution(st, event.Type)
		st.counter++
		st.totalSize += event.Size
	}

	return nil
}

func (h TrackStatsOnMediaImportedHandler) GetCurrentStats() StatsDto {
	result := make(chan StatsDto)
	h.ops <- func(st *statistics) {
		var avgSize int64
		if st.counter > 0 {
			avgSize = st.totalSize / st.counter
		}
		result <- StatsDto{st.counter, st.totalSize, avgSize, st.typeDistribution}
	}
	return <-result
}

func (h TrackStatsOnMediaImportedHandler) HandlerName() string {
	return "TrackStatsOnMediaImportedHandler"
}

func (h TrackStatsOnMediaImportedHandler) HandleGetStats(resp http.ResponseWriter, req *http.Request) {
	// get stats
	currStats := h.GetCurrentStats()

	// respond stats
	err := helper.RespondJSON(resp, currStats, nil)
	if err != nil {
		logrus.Errorf("could not marshall json for stats: %v", err)
		return
	}
}

func incTypeDistribution(st *statistics, t string) {
	currValue, _ := st.typeDistribution[t]
	st.typeDistribution[t] = currValue + 1
}

func (h TrackStatsOnMediaImportedHandler) run() {
	stats := statistics{
		typeDistribution: make(map[string]int),
	}
	for op := range h.ops {
		op(&stats)
	}
}
