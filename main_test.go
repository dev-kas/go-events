package events_test

import (
	"fmt"
	"testing"

	"glitchiethedev.com/events"
)

func Test(t *testing.T) {
	emitter := events.NewEventEmitter()

	listener := func(args ...interface{}) {
		fmt.Println("Event emitted with args:", args)
	}

	listenerID := emitter.On("test", listener)
	emitter.Emit("test", "arg1", "arg2", 3)
	emitter.RemoveListener("test", listenerID)
	emitter.Emit("test", "arg1", "arg2", 3)
}
