package main

import (
	"design-patterns-go/decoratorPattern"
	"fmt"
)

func main() {
	fmt.Println("Implementing basic design patterns in Go")
	fmt.Println(".")
	fmt.Println("..")
	fmt.Println("...")
	fmt.Println("....")
	fmt.Println("")

	//strategy := strategyPattern.NewStrategy()
	//strategy.Run()

	//observer := observerPattern.NewObserver()
	//observer.Run()

	decorator := decoratorPattern.NewDecorator()
	decorator.Run()
}
