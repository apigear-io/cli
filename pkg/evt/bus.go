package evt

type IHandler interface {
	HandleEvent(e *Event) (*Event, error)
}

type HandlerFunc func(req *Event) (*Event, error)

func (h HandlerFunc) HandleEvent(e *Event) (*Event, error) {
	return h(e)
}

type IEventBus interface {
	Publish(e *Event) error
	Register(kind string, fn HandlerFunc)
	Request(e *Event) (*Event, error)
	Use(mw ...HandlerFunc)
	Close() error
}
