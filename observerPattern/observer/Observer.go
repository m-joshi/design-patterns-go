package observer

import "design-patterns-go/observerPattern/event"

type Observer interface {
	OnNotify(event event.Event)
}
