package creational

import (
"fmt"
"time"
)

// Prototype Pattern
// Allows cloning of objects without coupling to their specific classes.
// Uses a prototype instance to create new objects by copying itself.

// Cloneable interface
type Prototype interface {
	Clone() Prototype
	GetInfo() string
}

// Document example
type Document struct {
	Title      string
	Content    string
	Author     string
	CreatedAt  time.Time
	ModifiedAt time.Time
	Tags       []string
	Metadata   map[string]string
}

func (d *Document) Clone() Prototype {
	// Deep copy
	tagsCopy := make([]string, len(d.Tags))
	copy(tagsCopy, d.Tags)

	metadataCopy := make(map[string]string)
	for k, v := range d.Metadata {
		metadataCopy[k] = v
	}

	return &Document{
		Title:      d.Title,
		Content:    d.Content,
		Author:     d.Author,
		CreatedAt:  d.CreatedAt,
		ModifiedAt: time.Now(),
		Tags:       tagsCopy,
		Metadata:   metadataCopy,
	}
}

func (d *Document) GetInfo() string {
	return fmt.Sprintf("Document: %s by %s (Created: %s)", d.Title, d.Author, d.CreatedAt.Format("2006-01-02"))
}

// Shape example with polymorphism
type Shape interface {
	Prototype
	Draw() string
}

type Circle struct {
	X      int
	Y      int
	Radius int
	Color  string
}

func (c *Circle) Clone() Prototype {
	return &Circle{
		X:      c.X,
		Y:      c.Y,
		Radius: c.Radius,
		Color:  c.Color,
	}
}

func (c *Circle) GetInfo() string {
	return fmt.Sprintf("Circle at (%d,%d) radius=%d color=%s", c.X, c.Y, c.Radius, c.Color)
}

func (c *Circle) Draw() string {
	return fmt.Sprintf("Drawing %s circle at (%d,%d) with radius %d", c.Color, c.X, c.Y, c.Radius)
}

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
	Color  string
}

func (r *Rectangle) Clone() Prototype {
	return &Rectangle{
		X:      r.X,
		Y:      r.Y,
		Width:  r.Width,
		Height: r.Height,
		Color:  r.Color,
	}
}

func (r *Rectangle) GetInfo() string {
	return fmt.Sprintf("Rectangle at (%d,%d) size=%dx%d color=%s", r.X, r.Y, r.Width, r.Height, r.Color)
}

func (r *Rectangle) Draw() string {
	return fmt.Sprintf("Drawing %s rectangle at (%d,%d) with size %dx%d", r.Color, r.X, r.Y, r.Width, r.Height)
}

// Prototype Registry - for managing prototypes
type PrototypeRegistry struct {
	prototypes map[string]Prototype
}

func NewPrototypeRegistry() *PrototypeRegistry {
	return &PrototypeRegistry{
		prototypes: make(map[string]Prototype),
	}
}

func (r *PrototypeRegistry) Register(name string, prototype Prototype) {
	r.prototypes[name] = prototype
}

func (r *PrototypeRegistry) Create(name string) (Prototype, error) {
	prototype, exists := r.prototypes[name]
	if !exists {
		return nil, fmt.Errorf("prototype '%s' not found", name)
	}
	return prototype.Clone(), nil
}

func (r *PrototypeRegistry) List() []string {
	names := make([]string, 0, len(r.prototypes))
	for name := range r.prototypes {
		names = append(names, name)
	}
	return names
}

// Real-world example: Database connection configuration
type DBConfig struct {
	Host            string
	Port            int
	Database        string
	Username        string
	Password        string
	MaxConnections  int
	Timeout         time.Duration
	SSL             bool
	ConnectionProps map[string]string
}

func (c *DBConfig) Clone() Prototype {
	propsCopy := make(map[string]string)
	for k, v := range c.ConnectionProps {
		propsCopy[k] = v
	}

	return &DBConfig{
		Host:            c.Host,
		Port:            c.Port,
		Database:        c.Database,
		Username:        c.Username,
		Password:        c.Password,
		MaxConnections:  c.MaxConnections,
		Timeout:         c.Timeout,
		SSL:             c.SSL,
		ConnectionProps: propsCopy,
	}
}

func (c *DBConfig) GetInfo() string {
	return fmt.Sprintf("DBConfig: %s@%s:%d/%s (SSL: %v)", c.Username, c.Host, c.Port, c.Database, c.SSL)
}

// Example usage demonstrating the pattern
func DemoPrototype() {
	fmt.Println("=== Prototype Pattern Demo ===\n")

	// Document cloning example
	fmt.Println("1. Document Cloning:")
	originalDoc := &Document{
		Title:      "Design Patterns",
		Content:    "This is a book about design patterns...",
		Author:     "Gang of Four",
		CreatedAt:  time.Now().Add(-24 * time.Hour),
		ModifiedAt: time.Now().Add(-24 * time.Hour),
		Tags:       []string{"programming", "design", "patterns"},
		Metadata:   map[string]string{"version": "1.0", "status": "published"},
	}

	fmt.Println("Original:", originalDoc.GetInfo())

	// Clone and modify
	clonedDoc := originalDoc.Clone().(*Document)
	clonedDoc.Title = "Design Patterns - Second Edition"
	clonedDoc.Tags = append(clonedDoc.Tags, "architecture")
	clonedDoc.Metadata["version"] = "2.0"

	fmt.Println("Cloned:  ", clonedDoc.GetInfo())
	fmt.Printf("Original tags: %v\n", originalDoc.Tags)
	fmt.Printf("Cloned tags:   %v\n", clonedDoc.Tags)
	fmt.Printf("Original version: %s\n", originalDoc.Metadata["version"])
	fmt.Printf("Cloned version:   %s\n", clonedDoc.Metadata["version"])

	// Shape cloning with registry
	fmt.Println("\n2. Shape Cloning with Registry:")
	registry := NewPrototypeRegistry()

	// Register prototypes
	registry.Register("red-circle", &Circle{X: 0, Y: 0, Radius: 10, Color: "red"})
	registry.Register("blue-rectangle", &Rectangle{X: 0, Y: 0, Width: 20, Height: 15, Color: "blue"})

	fmt.Println("Available prototypes:", registry.List())

	// Create instances from prototypes
	circle1, _ := registry.Create("red-circle")
	circle2, _ := registry.Create("red-circle")

	// Modify cloned circles
	c1 := circle1.(*Circle)
	c1.X = 10
	c1.Y = 20

	c2 := circle2.(*Circle)
	c2.X = 50
	c2.Y = 60
	c2.Radius = 25

	fmt.Println(c1.Draw())
	fmt.Println(c2.Draw())

	// Database configuration example
	fmt.Println("\n3. Database Configuration Templates:")

	// Create base configurations
	prodConfig := &DBConfig{
		Host:           "prod.example.com",
		Port:           5432,
		Database:       "myapp",
		Username:       "appuser",
		Password:       "prod_password",
		MaxConnections: 100,
		Timeout:        30 * time.Second,
		SSL:            true,
		ConnectionProps: map[string]string{
			"sslmode":     "require",
			"pool_size":   "20",
			"retry_limit": "3",
		},
	}

	// Clone for staging and development
	stagingConfig := prodConfig.Clone().(*DBConfig)
	stagingConfig.Host = "staging.example.com"
	stagingConfig.Password = "staging_password"
	stagingConfig.MaxConnections = 50
	stagingConfig.ConnectionProps["pool_size"] = "10"

	devConfig := prodConfig.Clone().(*DBConfig)
	devConfig.Host = "localhost"
	devConfig.Password = "dev_password"
	devConfig.MaxConnections = 10
	devConfig.SSL = false
	devConfig.ConnectionProps["sslmode"] = "disable"
	devConfig.ConnectionProps["pool_size"] = "5"

	fmt.Println("Production:", prodConfig.GetInfo())
	fmt.Println("Staging:   ", stagingConfig.GetInfo())
	fmt.Println("Development:", devConfig.GetInfo())

	// Verify deep copy
	fmt.Println("\n4. Verifying Deep Copy:")
	fmt.Printf("Prod pool_size:    %s\n", prodConfig.ConnectionProps["pool_size"])
	fmt.Printf("Staging pool_size: %s\n", stagingConfig.ConnectionProps["pool_size"])
	fmt.Printf("Dev pool_size:     %s\n", devConfig.ConnectionProps["pool_size"])
}
