package stats

import (
	"net/http"

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

type StatsStore struct {
	ops chan operation
}

func NewStatsStore() *StatsStore {
	store := &StatsStore{
		ops: make(chan operation),
	}
	http.HandleFunc("/stats", store.HandleGetStats)
	go store.run()
	return store
}

func (store *StatsStore) Track(mediaType string, mediaSize int64) {
	store.ops <- func(st *statistics) {
		st.incTypeDistribution(mediaType)
		st.counter++
		st.totalSize += mediaSize
	}
}

func (store StatsStore) GetCurrentStats() StatsDto {
	result := make(chan StatsDto)
	store.ops <- func(st *statistics) {
		var avgSize int64
		if st.counter > 0 {
			avgSize = st.totalSize / st.counter
		}
		result <- StatsDto{st.counter, st.totalSize, avgSize, st.typeDistribution}
	}
	return <-result
}

func (store StatsStore) HandleGetStats(resp http.ResponseWriter, req *http.Request) {
	// get stats
	currStats := store.GetCurrentStats()

	// respond stats
	err := helper.RespondJSON(resp, currStats, nil)
	if err != nil {
		logrus.Errorf("could not marshall json for stats: %v", err)
		return
	}
}

func (store *StatsStore) run() {
	stats := statistics{
		typeDistribution: make(map[string]int),
	}
	for op := range store.ops {
		op(&stats)
	}
}

func (st *statistics) incTypeDistribution(t string) {
	currValue, _ := st.typeDistribution[t]
	st.typeDistribution[t] = currValue + 1
}
