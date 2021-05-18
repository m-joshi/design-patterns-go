package duck

import "design-patterns-go/strategyPattern/flyingBehavior"

type RubberDuck struct {
	name string
}

func NewRubberDuck() RubberDuck {
	return RubberDuck{
		name:   "Rubber Duck",
	}
}

func (rd RubberDuck) Swim() string {
	return "I swim"
}

func (rd RubberDuck) Display() string  {
	return rd.name
}

func (rd RubberDuck) Fly() string {
	fly := flyingbehavior.FlyingBehavior{flyingbehavior.DontFly{}}
	return fly.PerformFly()
}

