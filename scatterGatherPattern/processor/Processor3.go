package processor

type Processor3 struct {

}


func NewProcessor3() Processor3 {
	return Processor3{}
}

func (p3 Processor3) Process() string {
	return "Processing 3"
}