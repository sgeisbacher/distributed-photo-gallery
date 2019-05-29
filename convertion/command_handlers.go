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

// GenerateThumbShotHandler command
type GenerateThumbShotHandler struct {
	eventBus *cqrs.EventBus
}

// NewCommand creates new GenerateThumbShot command
func (h GenerateThumbShotHandler) NewCommand() interface{} {
	// fmt.Println("GenerateThumbShotHandler NewCommand()")
	return &events.GenerateThumbShot{}
}

// Handle handles GenerateThumbShot command
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

// GenerateBigShotHandler command
type GenerateBigShotHandler struct {
	eventBus *cqrs.EventBus
}

// NewCommand creates new GenerateBigShot command
func (h GenerateBigShotHandler) NewCommand() interface{} {
	// fmt.Println("GenerateBigShotHandler NewCommand()")
	return &events.GenerateBigShot{}
}

// Handle handles GenerateBigShot command
func (h GenerateBigShotHandler) Handle(c interface{}) error {
	fmt.Println("generating big-shot ...")
	time.Sleep(2 * time.Second)
	fmt.Println("big-shot done")
	h.eventBus.Publish(context.Background(), events.MediaBigShotGenerated{})
	return nil
}
