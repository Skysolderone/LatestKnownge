package events

import "net"

type Event struct {
	Name string
	Data any
}
type NetEvent struct {
	Type    string
	Conn    net.Conn
	Message []byte
	Err     error
}
type EventListener func(NetEvent)

func NewEventListener(listener EventListener) EventListener {
	return listener
}

type Dispatcher struct {
	listeners map[string][]EventListener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[string][]EventListener),
	}
}
func (d *Dispatcher) RegisterListener(eventName string, listener EventListener) {
	d.listeners[eventName] = append(d.listeners[eventName], listener)
}
func (d *Dispatcher) DisPatch(event NetEvent) {
	for _, listen := range d.listeners[event.Type] {
		listen(event)
	}
}
func (d *Dispatcher) DisPatchAsync(event NetEvent) {
	for _, listen := range d.listeners[event.Type] {
		go listen(event)
	}
}
