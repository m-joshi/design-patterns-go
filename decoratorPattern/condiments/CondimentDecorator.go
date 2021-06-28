package condiments

type CondimentDecorator interface {
	GetDescription() string
	Cost() int
}
