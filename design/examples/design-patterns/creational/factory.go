package creational

// Factory Method - Creational Pattern
// Defines an interface for creating objects, but lets subclasses decide which class to instantiate

type Vehicle interface {
	Drive() string
}

type Car struct{}

func (c *Car) Drive() string {
	return "Driving a car"
}

type Bike struct{}

func (b *Bike) Drive() string {
	return "Riding a bike"
}

type VehicleFactory struct{}

func (f *VehicleFactory) CreateVehicle(vehicleType string) Vehicle {
	switch vehicleType {
	case "car":
		return &Car{}
	case "bike":
		return &Bike{}
	default:
		return nil
	}
}
