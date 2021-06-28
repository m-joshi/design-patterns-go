package beverage

type HouseBlend struct {
	description string
	cost        int
}

func NewHouseBlend() HouseBlend {
	return HouseBlend{
		description: "House Blend",
		cost:        50}
}

func (hb HouseBlend) GetDescription() string {
	return hb.description
}

func (hb HouseBlend) Cost() int {
	return hb.cost
}