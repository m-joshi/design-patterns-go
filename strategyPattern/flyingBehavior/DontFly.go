package flyingbehavior

type DontFly struct {}

func (DontFly) Fly() string {
	return "I don't fly"
}