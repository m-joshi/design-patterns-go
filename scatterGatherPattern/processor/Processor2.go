package processor

type Processor2 struct {

}


func NewProcessor2() Processor2 {
	return Processor2{}
}


func (p2 Processor2) Process() string {
	return "Processing 2"
}