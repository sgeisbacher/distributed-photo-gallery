package stats

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
)

func TestHandleMediaImported(t *testing.T) {
	RegisterTestingT(t)

	store := NewStatsStore()

	handler := TrackStatsOnMediaImportedHandler{nil, store}
	handler.Handle(context.Background(), &events.MediaImported{
		ID:   "1",
		Path: "/tmp/photos/1.jpg",
		Name: "1.jpg",
		Size: 550,
		Type: "jpg",
	})

	currStats := store.GetCurrentStats()
	Expect(currStats.Counter).To(Equal(int64(1)))
	Expect(currStats.TotalSize).To(Equal(int64(550)))
	Expect(currStats.AvgSize).To(Equal(int64(550)))
	Expect(len(currStats.TypeDistribution)).To(Equal(1))

	handler.Handle(context.Background(), &events.MediaImported{
		ID:   "2",
		Path: "/tmp/photos/2.jpg",
		Name: "2.jpg",
		Size: 750,
		Type: "jpg",
	})
	handler.Handle(context.Background(), &events.MediaImported{
		ID:   "3",
		Path: "/tmp/photos/3.png",
		Name: "3.png",
		Size: 100,
		Type: "png",
	})

	currStats = store.GetCurrentStats()
	Expect(currStats.Counter).To(Equal(int64(3)))
	Expect(currStats.TotalSize).To(Equal(int64(1400)))
	Expect(currStats.AvgSize).To(Equal(int64(466)))
	Expect(len(currStats.TypeDistribution)).To(Equal(2))
}
