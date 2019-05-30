package convertion

import (
	"context"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sgeisbacher/distributed-photo-gallery/events"
)

/*
*
*      GenerateThumbShot-command-handler
*
 */

type GenerateThumbShotHandler struct {
	eventBus *cqrs.EventBus
}

func (h GenerateThumbShotHandler) NewCommand() interface{} {
	return &events.GenerateThumbShot{}
}

func (h GenerateThumbShotHandler) Handle(c interface{}) error {
	fmt.Println("generating thumb-shot ...")
	time.Sleep(2 * time.Second)
	fmt.Println("thumb-shot done")
	h.eventBus.Publish(context.Background(), events.MediaThumbShotGenerated{})
	return nil
}

/*
*
*      GenerateBigShot-command-handler
*
 */

type GenerateBigShotHandler struct {
	eventBus *cqrs.EventBus
}

func (h GenerateBigShotHandler) NewCommand() interface{} {
	return &events.GenerateBigShot{}
}

func (h GenerateBigShotHandler) Handle(c interface{}) error {
	fmt.Println("generating big-shot ...")
	time.Sleep(2 * time.Second)
	fmt.Println("big-shot done")
	h.eventBus.Publish(context.Background(), events.MediaBigShotGenerated{})
	return nil
}
