package event

type Weather struct {
	temperature int
	humidity    int
	pressure    int
}

func NewWeather(temperature int, humidity int, pressure int) Weather {
	return Weather{
		temperature: temperature,
		humidity:    humidity,
		pressure:    pressure,
	}
}

func (w Weather) Display() {
	println("Temperature : ", w.temperature)
	println("Humidity : ", w.humidity)
	println("Pressure : ", w.pressure)
}
