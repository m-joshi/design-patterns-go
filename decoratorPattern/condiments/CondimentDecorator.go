package condiments

type CondimentDecorator interface {
	GetDescription(func() string) func() string
	Cost(func() int) func() int
}
