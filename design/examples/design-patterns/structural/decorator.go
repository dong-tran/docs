package structural

// Decorator - Structural Pattern
// Adds new functionality to objects dynamically

type Coffee interface {
	Cost() float64
	Description() string
}

type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() float64 {
	return 2.0
}

func (c *SimpleCoffee) Description() string {
	return "Simple coffee"
}

// Decorators
type MilkDecorator struct {
	coffee Coffee
}

func (m *MilkDecorator) Cost() float64 {
	return m.coffee.Cost() + 0.5
}

func (m *MilkDecorator) Description() string {
	return m.coffee.Description() + ", milk"
}

type SugarDecorator struct {
	coffee Coffee
}

func (s *SugarDecorator) Cost() float64 {
	return s.coffee.Cost() + 0.2
}

func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + ", sugar"
}

// Usage: coffee := &SugarDecorator{&MilkDecorator{&SimpleCoffee{}}}
