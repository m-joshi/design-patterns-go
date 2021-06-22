package notifier

import (
	"design-patterns-go/observerPattern/event"
	"design-patterns-go/observerPattern/observer"
)

type WeatherStation struct {
	weather event.Weather
	observers map[observer.Observer]struct{}
}

func NewWeatherStation(weather event.Weather) WeatherStation {
	return WeatherStation{
		weather: weather,
		observers: make(map[observer.Observer]struct{}),
	}
}

func (ws WeatherStation) Register(observer observer.Observer)  {
	println("Registering Observer")
	ws.observers[observer] = struct{}{}
}

func (ws WeatherStation) Deregister(observer observer.Observer)  {
	println("Deregistering Observer")
	delete(ws.observers, observer)
}

func (ws WeatherStation) Notify(event event.Event)  {
	for observer := range ws.observers {
		observer.OnNotify(event)
	}
}

func (ws WeatherStation) ChangeWeather(weather event.Weather)  {
	ws.weather = weather
	var event event.Event = weather
	ws.Notify(event)
}