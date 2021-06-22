package observerPattern

import (
	"design-patterns-go/observerPattern/event"
	"design-patterns-go/observerPattern/notifier"
	"design-patterns-go/observerPattern/observer"
	"fmt"
)

type ObserverPattern interface {
	Run()
}

type observerPattern struct {
	observerPattern ObserverPattern
}

func NewObserver() observerPattern {
	return observerPattern{}
}

func (ob observerPattern) Run() {
	fmt.Println("~~~~~~~~~~ OBSERVER PATTERN ~~~~~~~~~~~~~")

	fmt.Println("")
	fmt.Println("Initialising weather station with initial weather")
	weather := event.NewWeather(26, 85, 1014)
	weatherStation := notifier.NewWeatherStation(weather)

	fmt.Println("-------------------------------------------------")
	fmt.Println("Adding observers")

	boringAppDisplay := observer.NewBoringAppDisplay()
	weatherStation.Register(boringAppDisplay)


	billBoardDisplay := observer.NewBillBoardDisplay()
	weatherStation.Register(billBoardDisplay)

	fmt.Println("")
	fmt.Println("Weather Changing due to thunderstorm")
	thunderstormEvent := event.NewWeather(18, 92, 1134)
	weatherStation.ChangeWeather(thunderstormEvent)

	fmt.Println("")
	fmt.Println("Can stand boring app no more")
	weatherStation.Deregister(boringAppDisplay)

	fmt.Println("Weather change due to scorching heat")
	scorchingHeatEvent := event.NewWeather(40, 60, 934)
	weatherStation.ChangeWeather(scorchingHeatEvent)

}
