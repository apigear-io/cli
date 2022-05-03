package mon

type EventSender struct {
	addr string
}

func NewEventSender(addr string) *EventSender {
	return &EventSender{addr: addr}
}

// SendEvents sends events to the monitor server.
// The events are sent as json encoded strings.
// The events are sent to the monitor server using a http post message

func (s *EventSender) SendEvents(emitter chan *Event) {
	for event := range emitter {
		go HttpPost(s.addr, event)
	}
}
