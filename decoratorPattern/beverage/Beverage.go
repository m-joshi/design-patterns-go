package beverage

type Beverage interface {
	GetDescription() string
	Cost() int
}