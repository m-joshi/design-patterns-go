package notifier

import (
	"design-patterns-go/observerPattern/event"
	"design-patterns-go/observerPattern/observer"
)

type Notifier interface {
	Register(observer observer.Observer)
	Deregister(observer observer.Observer)
	Notify(event event.Event)
}
