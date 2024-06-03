# Go-Events
A [`node:events`](https://nodejs.org/api/events.html) replica from NodeJs

---

Install via
```sh
go get github.com/dev-kas/go-events
```

Import it via
```go
import "github.com/dev-kas/go-events"
// or
import (
    ...
    "github.com/dev-kas/go-events"
    ...
)
```

Create a new `EventEmitter` via
```go
emitter := events.NewEventEmitter()
// or
var emitter *events.EventEmitter
emitter = events.NewEventEmitter()
```

Listen for events via
```go
listenerID := emitter.On("my event", func(args ...interface{}) {
    fmt.Println("Event emitted with args:", args)
})
```

Emit events via
```go
emitter.Emit("my event", "arg1", "arg2", 3)
```

Remove event listeners via
```go
emitter.RemoveListener("my event", listenerID)
```

See `main_test.go`.

It's pretty much same as [`node:events`](https://nodejs.org/api/events.html).
