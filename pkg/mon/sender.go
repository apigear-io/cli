package mon

import (
	"time"
)

// EventSender is a sender of events.
// It sends events to the monitor server
type EventSender struct {
	url string
}

func NewEventSender(url string) *EventSender {
	return &EventSender{url: url}
}

// SendEvents sends events to the monitor server.
// The events are sent as json encoded strings.
// The events are sent to the monitor server using a http post message
func (s *EventSender) SendEvents(events []*Event, sleep time.Duration) {
	for _, event := range events {
		log.Info().Msgf("send event: %+v", event)
		// capture url, event for closure
		err := HttpPost(s.url, event)
		if err != nil {
			log.Warn().Msgf("send event: %s", err)
		}
		if sleep > 0 {
			time.Sleep(sleep)
		}
	}
}
