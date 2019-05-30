package store

import (
	"context"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
	"github.com/sgeisbacher/distributed-photo-gallery/helper"
	"github.com/sirupsen/logrus"
)

type operation func(map[string]Media)

type CreateMediaOnMediaImportedHandler struct {
	CommandBus *cqrs.CommandBus
	ops        chan operation
}

func NewCreateMediaOnMediaImportedHandler(cb *cqrs.CommandBus) CreateMediaOnMediaImportedHandler {
	handler := CreateMediaOnMediaImportedHandler{
		CommandBus: cb,
		ops:        make(chan operation),
	}
	http.HandleFunc("/media", handler.HandleGetAllMedias)
	go handler.run()
	return handler
}

func (h CreateMediaOnMediaImportedHandler) NewEvent() interface{} {
	return &events.MediaImported{}
}

func (h CreateMediaOnMediaImportedHandler) Handle(ctx context.Context, e interface{}) error {
	event := e.(*events.MediaImported)
	logrus.Debugf("creating media %q in db", event.ID)
	h.ops <- func(db map[string]Media) {
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
	}
	return nil
}

func (h CreateMediaOnMediaImportedHandler) HandlerName() string {
	return "CreateMediaOnMediaImportedHandler"
}

func (h CreateMediaOnMediaImportedHandler) HandleGetAllMedias(resp http.ResponseWriter, req *http.Request) {
	result := make(chan []Media)
	h.ops <- func(db map[string]Media) {
		var medias []Media
		for _, media := range db {
			medias = append(medias, media)
		}
		result <- medias
	}
	medias := <-result
	err := helper.RespondJSON(resp, medias, nil)
	if err != nil {
		logrus.Errorf("error while responding GetAllMedias: %v", err)
	}
}

func (h CreateMediaOnMediaImportedHandler) run() {
	db := make(map[string]Media)
	for op := range h.ops {
		op(db)
	}
}
