package stats

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
	"github.com/sirupsen/logrus"
)

var counter int64
var totalSize int64
var typeDistribution map[string]int

func init() {
	typeDistribution = make(map[string]int)
	http.HandleFunc("/stats", HandleGetStats)
}

// TrackStatsOnMediaImportedHandler event handler
type TrackStatsOnMediaImportedHandler struct {
	CommandBus *cqrs.CommandBus
}

// NewEvent creates new event
func (h TrackStatsOnMediaImportedHandler) NewEvent() interface{} {
	return &events.MediaImported{}
}

// Handle handle event
func (h TrackStatsOnMediaImportedHandler) Handle(e interface{}) error {
	event := e.(*events.MediaImported)
	logrus.Debug("tracking stats for media:", event.Path)

	// track
	incTypeDistribution(event.Type)
	counter++
	totalSize += event.Size

	return nil
}

func HandleGetStats(resp http.ResponseWriter, req *http.Request) {
	if counter < 1 {
		resp.WriteHeader(http.StatusNoContent)
		return
	}
	resp.Header().Add("Content-Type", "application/json")
	jsonData, err := json.Marshal(struct {
		Counter          int64
		TotalSize        int64
		AvgSize          int64
		TypeDistribution map[string]int
	}{counter, totalSize, totalSize / counter, typeDistribution})
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("could not marshall json for stats: %v", err)
		return
	}
	fmt.Fprintln(resp, string(jsonData))
}

func incTypeDistribution(t string) {
	currValue, _ := typeDistribution[t]
	typeDistribution[t] = currValue + 1
}
