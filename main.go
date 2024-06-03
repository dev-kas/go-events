package events

import (
    "sync"
)

type event struct {
    listeners map[int]func(...interface{})
    idCounter int
    lock      sync.Mutex
}

func newEvent() *event {
    return &event{
        listeners: make(map[int]func(...interface{})),
    }
}

func (e *event) addListener(listener func(...interface{})) int {
    e.lock.Lock()
    defer e.lock.Unlock()
    e.idCounter++
    e.listeners[e.idCounter] = listener
    return e.idCounter
}

func (e *event) removeListener(id int) {
    e.lock.Lock()
    defer e.lock.Unlock()
    delete(e.listeners, id)
}

type EventEmitter struct {
    events map[string]*event
    lock   sync.Mutex
}

func NewEventEmitter() *EventEmitter {
    return &EventEmitter{
        events: make(map[string]*event),
    }
}

func (emitter *EventEmitter) On(event string, listener func(...interface{})) int {
    emitter.lock.Lock()
    defer emitter.lock.Unlock()
    if _, ok := emitter.events[event]; !ok {
        emitter.events[event] = newEvent()
    }
    return emitter.events[event].addListener(listener)
}

func (emitter *EventEmitter) Emit(event string, args ...interface{}) {
    emitter.lock.Lock()
    defer emitter.lock.Unlock()
    if e, ok := emitter.events[event]; ok {
        for _, listener := range e.listeners {
            listener(args...)
        }
    }
}

func (emitter *EventEmitter) RemoveListener(event string, id int) {
    emitter.lock.Lock()
    defer emitter.lock.Unlock()
    if e, ok := emitter.events[event]; ok {
        e.removeListener(id)
    }
}

func (emitter *EventEmitter) RemoveAllListeners(event string) {
    emitter.lock.Lock()
    defer emitter.lock.Unlock()
    delete(emitter.events, event)
}
