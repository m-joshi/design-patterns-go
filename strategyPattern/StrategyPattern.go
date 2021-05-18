package strategyPattern

import (
	"design-patterns-go/strategyPattern/duck"
	"fmt"
)

type StrategyPattern interface {
	Run()
}

type strategy struct {
	strategyPattern StrategyPattern
}

func NewStrategy() strategy {
	return strategy{}
}

func (st strategy) Run() {
	fmt.Println("~~~~~~~~~~ STRATEGY PATTERN ~~~~~~~~~~~~~")

	mallardDuck := duck.NewMallardDuck()
	fmt.Println("name : ", mallardDuck.Display())
	fmt.Println("swim : ", mallardDuck.Swim())
	fmt.Println("fly : ", mallardDuck.Fly())

	rubberDuck := duck.NewRubberDuck()
	fmt.Println("name : ", rubberDuck.Display())
	fmt.Println("swim : ", rubberDuck.Swim())
	fmt.Println("fly : ", rubberDuck.Fly())
}