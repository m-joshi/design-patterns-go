package observer

import "design-patterns-go/observerPattern/event"

type BoringAppDisplay struct {

}

func NewBoringAppDisplay() BoringAppDisplay {
	println("Creating Boring App Display Observer")
	return BoringAppDisplay{}
}

func (bad BoringAppDisplay) OnNotify(event event.Event)  {
	println("~~~~~~~Boring App Display~~~~~~~~~~")
	event.Display()
}