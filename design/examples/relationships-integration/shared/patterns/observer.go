package patterns

// Observer Pattern - notifies multiple subscribers of events
type Event struct {
	Type string
	Data interface{}
}

type EventObserver interface {
	OnEvent(event Event)
}

type EventPublisher struct {
	observers []EventObserver
}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{observers: make([]EventObserver, 0)}
}

func (p *EventPublisher) Subscribe(observer EventObserver) {
	p.observers = append(p.observers, observer)
}

func (p *EventPublisher) Publish(event Event) {
	for _, observer := range p.observers {
		observer.OnEvent(event)
	}
}
