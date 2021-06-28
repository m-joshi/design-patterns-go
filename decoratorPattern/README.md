# Decorator Pattern

- Attaches additional responsibilities to an object dynamically. 
- Decorators provide a flexible alternative to subclassing for extending functionality.
- Applies SOLID Principle - Open for extension, closed for modification.

### Java Way

- Going by the book, we wrapped the main component in the Decorator component.
```
type Component interface {
    Cost() int
}

type Decorator interface {
    Cost() int 
}

type Decorator struct {
    component Component
    cost int
}

func (dc Decorator) Cost() int {
    dc.component.Cost() + dc.cost 
}
```

### Pure Go Way

- In Go, functions can be :
    - Can be passed around as arguments in other functions
    - Can be assigned to variables
    - Can be returned by functions
- We leverage this in implementing Decorator pattern
- Function of the main Component is passed as value in the function of Decorator component
```
type Component interface {
    Cost() int
}

type Decorator interface {
    Cost(func() int) func() int 
}

type Decorator struct {
    cost int
}

func (dc Decorator) Cost(componentCost func() int) func() int {
    return func() int {
        return componentCost() + dc.cost
    } 
}

func main() {
    finalCost := (Decorator.Cost(Component.Cost))()
}
```
 