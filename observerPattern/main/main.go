package main

import (
	"design-patterns-go/observerPattern"
	"fmt"
)

func main() {
	fmt.Println("Implementing Observer Pattern in Go")
	fmt.Println(".")
	fmt.Println("..")
	fmt.Println("...")
	fmt.Println("....")
	fmt.Println("")

	observer := observerPattern.NewObserver()
	observer.Run()
}
