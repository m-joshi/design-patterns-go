package main

import (
	"design-patterns-go/strategyPattern"
	"fmt"
)

func main() {
	fmt.Println("Implementing basic design patterns in Go")

	strategy := strategyPattern.NewStrategy()
	strategy.Run()
}
