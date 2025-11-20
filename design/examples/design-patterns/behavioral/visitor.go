package behavioral

import "fmt"

// Visitor Pattern - Allows adding new operations to objects without modifying them.

type Visitor interface {
	VisitCircle(*Circle) string
	VisitRectangle(*Rectangle) string
	VisitTriangle(*Triangle) string
}

type Shape interface {
	Accept(Visitor) string
}

type Circle struct {
	Radius float64
}

func (c *Circle) Accept(v Visitor) string {
	return v.VisitCircle(c)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(v Visitor) string {
	return v.VisitRectangle(r)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) Accept(v Visitor) string {
	return v.VisitTriangle(t)
}

type AreaCalculator struct{}

func (a *AreaCalculator) VisitCircle(c *Circle) string {
	area := 3.14159 * c.Radius * c.Radius
	return fmt.Sprintf("Circle area: %.2f", area)
}

func (a *AreaCalculator) VisitRectangle(r *Rectangle) string {
	area := r.Width * r.Height
	return fmt.Sprintf("Rectangle area: %.2f", area)
}

func (a *AreaCalculator) VisitTriangle(t *Triangle) string {
	area := 0.5 * t.Base * t.Height
	return fmt.Sprintf("Triangle area: %.2f", area)
}

type PerimeterCalculator struct{}

func (p *PerimeterCalculator) VisitCircle(c *Circle) string {
	perimeter := 2 * 3.14159 * c.Radius
	return fmt.Sprintf("Circle perimeter: %.2f", perimeter)
}

func (p *PerimeterCalculator) VisitRectangle(r *Rectangle) string {
	perimeter := 2 * (r.Width + r.Height)
	return fmt.Sprintf("Rectangle perimeter: %.2f", perimeter)
}

func (p *PerimeterCalculator) VisitTriangle(t *Triangle) string {
	perimeter := 3 * t.Base
	return fmt.Sprintf("Triangle perimeter: %.2f", perimeter)
}

type JSONExporter struct{}

func (j *JSONExporter) VisitCircle(c *Circle) string {
	return fmt.Sprintf(`{"type": "circle", "radius": %.2f}`, c.Radius)
}

func (j *JSONExporter) VisitRectangle(r *Rectangle) string {
	return fmt.Sprintf(`{"type": "rectangle", "width": %.2f, "height": %.2f}`, r.Width, r.Height)
}

func (j *JSONExporter) VisitTriangle(t *Triangle) string {
	return fmt.Sprintf(`{"type": "triangle", "base": %.2f, "height": %.2f}`, t.Base, t.Height)
}

func DemoVisitor() {
	fmt.Println("=== Visitor Pattern Demo ===\n")
	shapes := []Shape{
		&Circle{Radius: 5},
		&Rectangle{Width: 4, Height: 6},
		&Triangle{Base: 3, Height: 4},
	}
	areaCalc := &AreaCalculator{}
	perimeterCalc := &PerimeterCalculator{}
	jsonExporter := &JSONExporter{}
	fmt.Println("Calculating Areas:")
	for _, shape := range shapes {
		fmt.Println(shape.Accept(areaCalc))
	}
	fmt.Println("\nCalculating Perimeters:")
	for _, shape := range shapes {
		fmt.Println(shape.Accept(perimeterCalc))
	}
	fmt.Println("\nExporting to JSON:")
	for _, shape := range shapes {
		fmt.Println(shape.Accept(jsonExporter))
	}
}
