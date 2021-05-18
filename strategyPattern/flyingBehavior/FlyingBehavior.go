package flyingbehavior

type Flying interface {
	Fly() string
}

type FlyingBehavior struct {
	Flying Flying
}

func (fb FlyingBehavior) PerformFly() string {
	return fb.Flying.Fly()
}