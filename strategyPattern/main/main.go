package main

import (
	"design-patterns-go/strategyPattern"
	"fmt"
)

func main() {
	fmt.Println("Implementing Strategy Pattern in Go")
	fmt.Println(".")
	fmt.Println("..")
	fmt.Println("...")
	fmt.Println("....")
	fmt.Println("")

	strategy := strategyPattern.NewStrategy()
	strategy.Run()
}
