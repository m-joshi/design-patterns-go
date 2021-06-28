package condiments

type Mocha struct {
	description string
	cost        int
}

func NewMocha() Mocha {
	return Mocha{
		description: "Mocha",
		cost:        20,
	}
}

func (mc Mocha)GetDescription(beverageDescription func() string) func() string {
	return func() string {
		return beverageDescription() +" "+ mc.description
	}
}

func (mc Mocha)Cost(beverageCost func() int) func() int {
	return func() int {
		return  beverageCost() + mc.cost
	}
}
