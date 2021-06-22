package observer

import "design-patterns-go/observerPattern/event"

type BillBoardDisplay struct {

}

func NewBillBoardDisplay() BillBoardDisplay {
	println("Creating Bill Board Display Observer")

	return BillBoardDisplay{}
}

func (bbd BillBoardDisplay) OnNotify(event event.Event)  {
	println("~~~~~~~Bill Board Display~~~~~~~~~~")
	event.Display()
}