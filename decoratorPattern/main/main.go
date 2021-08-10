package main

import (
	"design-patterns-go/decoratorPattern"
	"fmt"
)

func main() {
	fmt.Println("Implementing Decorator Pattern in Go")
	fmt.Println(".")
	fmt.Println("..")
	fmt.Println("...")
	fmt.Println("....")
	fmt.Println("")

	decorator := decoratorPattern.NewDecorator()
	decorator.Run()
}
