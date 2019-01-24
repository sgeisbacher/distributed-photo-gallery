package stats

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
)

var counter int
var typeDistribution map[string]int

func init() {
	typeDistribution = make(map[string]int)
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
	fmt.Println("tracking stats for media:", event.Path)

	// track
	incTypeDistribution(event.Type)
	counter++

	return nil
}

// Print prints current stats to console
func Print() {
	fmt.Println("stats:")
	fmt.Printf("\tcounter: %d\n", counter)
	fmt.Printf("\ttype-distribution: %v\n", typeDistribution)
}

func incTypeDistribution(t string) {
	currValue, _ := typeDistribution[t]
	typeDistribution[t] = currValue + 1
}
