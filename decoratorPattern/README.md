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
