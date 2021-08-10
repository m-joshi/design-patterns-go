package main

import (
	"design-patterns-go/scatterGatherPattern"
	"fmt"
)

func main() {
	fmt.Println("Implementing Scatter Gather Pattern in Go")
	fmt.Println(".")
	fmt.Println("..")
	fmt.Println("...")
	fmt.Println("....")
	fmt.Println("")

	scatterGather := scatterGatherPattern.NewScatterGather()
	scatterGather.Run()
}
