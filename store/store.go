package store

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
	"github.com/sirupsen/logrus"
)

// todo var preStage map[string]Media
var db map[string]Media

func init() {
	db = make(map[string]Media)
	http.HandleFunc("/media", HandleGetAllMedias)
}

// CreateMediaOnMediaImportedHandler event handler
type CreateMediaOnMediaImportedHandler struct {
	CommandBus *cqrs.CommandBus
}

// NewEvent creates new event
func (h CreateMediaOnMediaImportedHandler) NewEvent() interface{} {
	return &events.MediaImported{}
}

// Handle handle event
func (h CreateMediaOnMediaImportedHandler) Handle(ctx context.Context, e interface{}) error {
	// todo race
	event := e.(*events.MediaImported)
	logrus.Debugf("creating media %q in db", event.ID)
	if _, ok := db[event.ID]; ok {
		logrus.Warn("already exists, overriding ...")
	}
	db[event.ID] = Media{
		ID:       event.ID,
		Name:     event.Name,
		OrigPath: event.Path,
		Type:     event.Type,
		Status:   created,
	}
	return nil
}

func (h CreateMediaOnMediaImportedHandler) HandlerName() string {
	return "CreateMediaOnMediaImportedHandler"
}

func HandleGetAllMedias(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "application/json")
	jsondata, err := json.Marshal(db)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("could not marshall json for all medias: %v", err)
		return
	}
	fmt.Fprintln(resp, string(jsondata))
}
