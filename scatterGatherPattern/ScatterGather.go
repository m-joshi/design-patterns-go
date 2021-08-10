package scatterGatherPattern

import (
	"design-patterns-go/scatterGatherPattern/processor"
	"fmt"
)

type ScatterGather interface {
	Run()
}

type scatterGather struct {
	scatterGather ScatterGather
	processors []processor.Processor
	results chan string
}

func NewScatterGather() scatterGather {
	return scatterGather{}
}

func (sg *scatterGather) Run() {
	sg.createProcessors()

	sg.scatter()

	sg.gather()
}

func (sg *scatterGather) gather() {
	for range sg.processors {
		result := <- sg.results
		fmt.Println(result)
	}
}

func (sg *scatterGather) scatter() {
	sg.results = make(chan string, 0)
	for _, p := range sg.processors {
		go func(pro processor.Processor) {
			result := pro.Process()
			sg.results <- result
		}(p)
	}
}

func (sg *scatterGather) createProcessors() {
	processor1 := processor.NewProcessor1()
	processor2 := processor.NewProcessor2()
	processor3 := processor.NewProcessor3()

	sg.processors = append(sg.processors, processor1, processor2, processor3)
}

