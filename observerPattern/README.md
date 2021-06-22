# Observer Pattern

- Defines one-to-many dependency between objects
- When the state of one object change, all its dependent objects are notified.

### Example Taken
A weather station records change in the weather and notifies two displays.

### Three major components

#### Event
- Captures change in the state
```
type Event struct {
    Data int
}
```

#### Observers
- Gets notified whenever there is a change in the notifier it is registered to.
- Different observers can behave differently when they are notified.
```
type Observer interface {
	OnNotify(event event.Event)
}
```
<em>In our example, we have two displays as observers. They implement `OnNotify` function differently.</em>

#### Notifier
- This is the object that changes 
- It has a list of Observers
- Supports registering, deregistering and notifying the observers
```
type struct Notifier struct {
    observers map[observer.Observer]struct{}
    Data int
}

type Notifier interface {
	Register(observer observer.Observer)
	Deregister(observer observer.Observer)
	Notify(event event.Event)
}
```
<em>In our example, Weather station is a Notifier. It implements the functions of `Notifier` interface.</em>

#### Crucial Stuff
- `Notify` function of Notifier calls `OnNotify` function of all observers
- `Notify` function itself can be triggered at the change of the state
```
func Notify(event event.Event)  {
	for observer := range observers {
		observer.OnNotify(event)
	}
}

func ChangeState(event event.Event)  {
    // actually change state
	Notify(event)
}
```


