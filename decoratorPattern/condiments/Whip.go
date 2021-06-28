package condiments

type Whip struct {
	description string
	cost        int
}

func NewWhip() Whip {
	return Whip{
		description: "Whip",
		cost:        30,
	}
}

func (wp Whip)GetDescription(beverageDescription func() string) func() string {
	return func() string {
		return beverageDescription() +" "+ wp.description
	}
}

func (wp Whip)Cost(beverageCost func() int) func() int {
	return func() int {
		return  beverageCost() + wp.cost
	}
}
