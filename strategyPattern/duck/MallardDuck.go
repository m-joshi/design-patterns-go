package duck

import "design-patterns-go/strategyPattern/flyingBehavior"

type MallardDuck struct {
	name string
}

func NewMallardDuck() MallardDuck {
	return MallardDuck{
		name:   "Mallard Duck",
	}
}

func (md MallardDuck) Swim() string {
	return "I swim"
}

func (md MallardDuck) Display() string  {
	return md.name
}

func (md MallardDuck) Fly() string {
	fly := flyingbehavior.FlyingBehavior{flyingbehavior.FlyWithWings{}}
	return fly.PerformFly()
}
