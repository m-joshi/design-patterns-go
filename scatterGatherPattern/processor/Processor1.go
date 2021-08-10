package processor

type Processor1 struct {

}

func NewProcessor1() Processor1 {
	return Processor1{}
}

func (p1 Processor1) Process() string {
	return "Processing 1"
}