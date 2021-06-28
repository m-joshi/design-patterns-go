# design-patterns-go

- This repository has design patterns from the book ["Head First Design Patterns"](https://www.oreilly.com/library/view/head-first-design/0596007124/)
- The aim is to implement OO concepts in Go as cleanly and effortlessly as possible. 

### Build
- Since there are various patterns, make appropriate changes in `main.go`
- Comment all patterns, other than what you want to run

```
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
```

- Change `BINARY_NAME` in `Makefile`
- Run task to build binary
```
$make build
```

- Add a make task to run the binary
```
run_decorator:
	./decorator_pattern
```
- Run task
```
$make run_decorator
```