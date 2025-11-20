package behavioral

// Observer - Behavioral Pattern
// Defines one-to-many dependency between objects

type Observer interface {
	Update(temp float64)
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

type WeatherStation struct {
	observers   []Observer
	temperature float64
}

func (w *WeatherStation) Attach(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *WeatherStation) Detach(observer Observer) {
	// Remove observer
}

func (w *WeatherStation) Notify() {
	for _, observer := range w.observers {
		observer.Update(w.temperature)
	}
}

func (w *WeatherStation) SetTemperature(temp float64) {
	w.temperature = temp
	w.Notify()
}

type PhoneDisplay struct {
	name string
}

func (p *PhoneDisplay) Update(temp float64) {
	// Update phone display
}

type WebDisplay struct {
	name string
}

func (w *WebDisplay) Update(temp float64) {
	// Update web display
}
