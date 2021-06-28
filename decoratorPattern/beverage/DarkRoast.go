package beverage

type DarkRoast struct {
	description string
	cost int
}


func NewDarkRoast() DarkRoast {
	return DarkRoast{
		description: "Dark Roast",
		cost:        60}
}

func (dr DarkRoast) GetDescription() string {
	return dr.description
}

func (dr DarkRoast) Cost() int {
	return dr.cost
}