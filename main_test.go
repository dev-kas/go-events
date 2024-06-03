package events_test

import (
	"fmt"
	"testing"

	"github.com/dev-kas/go-events"
)

func Test(t *testing.T) {
	emitter := events.NewEventEmitter()

	listenerID := emitter.On("my event", func(args ...interface{}) {
		fmt.Println("Event emitted with args:", args)
	})

	emitter.Emit("my event", "arg1", "arg2", 3)

	emitter.RemoveListener("my event", listenerID)

	emitter.Emit("my event", "arg1", "arg2", 3)
}
