package condiments

type Milk struct {
	description string
	cost        int
}

func NewMilk() Milk {
	return Milk{
		description: "Milk",
		cost:        10,
	}
}

func (ml Milk)GetDescription(beverageDescription func() string) func() string {
	return func() string {
		return beverageDescription() +" "+ ml.description
	}
}

func (ml Milk)Cost(beverageCost func() int) func() int {
	return func() int {
		return  beverageCost() + ml.cost
	}
}
