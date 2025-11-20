package structural

import "fmt"

// Composite Pattern
// Composes objects into tree structures to represent part-whole hierarchies.
// Allows clients to treat individual objects and compositions uniformly.

type Component interface {
	Operation() string
	Add(Component)
	Remove(Component)
	GetChild(int) Component
}

// Leaf
type File struct {
	name string
}

func (f *File) Operation() string {
	return f.name
}

func (f *File) Add(Component)       { /* Files can't have children */ }
func (f *File) Remove(Component)    { /* Files can't have children */ }
func (f *File) GetChild(int) Component { return nil }

// Composite
type Folder struct {
	name     string
	children []Component
}

func (f *Folder) Operation() string {
	result := f.name + "\n"
	for _, child := range f.children {
		result += "  " + child.Operation() + "\n"
	}
	return result
}

func (f *Folder) Add(c Component) {
	f.children = append(f.children, c)
}

func (f *Folder) Remove(c Component) {
	for i, child := range f.children {
		if child == c {
			f.children = append(f.children[:i], f.children[i+1:]...)
			break
		}
	}
}

func (f *Folder) GetChild(index int) Component {
	if index >= 0 && index < len(f.children) {
		return f.children[index]
	}
	return nil
}

// Real-world example: Graphics system
type Graphic interface {
	Draw() string
	Move(x, y int)
}

type Dot struct {
	x, y int
}

func (d *Dot) Draw() string {
	return fmt.Sprintf("Dot at (%d, %d)", d.x, d.y)
}

func (d *Dot) Move(x, y int) {
	d.x += x
	d.y += y
}

type Circle struct {
	Dot
	radius int
}

func (c *Circle) Draw() string {
	return fmt.Sprintf("Circle at (%d, %d) with radius %d", c.x, c.y, c.radius)
}

type CompoundGraphic struct {
	children []Graphic
}

func (cg *CompoundGraphic) Draw() string {
	result := "Compound Graphic:\n"
	for _, child := range cg.children {
		result += "  " + child.Draw() + "\n"
	}
	return result
}

func (cg *CompoundGraphic) Move(x, y int) {
	for _, child := range cg.children {
		child.Move(x, y)
	}
}

func (cg *CompoundGraphic) Add(g Graphic) {
	cg.children = append(cg.children, g)
}

func DemoComposite() {
	fmt.Println("=== Composite Pattern Demo ===\n")

	// File system example
	fmt.Println("1. File System:")
	root := &Folder{name: "/"}
	home := &Folder{name: "home"}
	docs := &Folder{name: "documents"}

	file1 := &File{name: "resume.pdf"}
	file2 := &File{name: "photo.jpg"}
	file3 := &File{name: "config.txt"}

	docs.Add(file1)
	docs.Add(file2)
	home.Add(docs)
	home.Add(file3)
	root.Add(home)

	fmt.Print(root.Operation())

	// Graphics example
	fmt.Println("\n2. Graphics System:")
	all := &CompoundGraphic{}

	dot1 := &Dot{x: 1, y: 2}
	dot2 := &Dot{x: 5, y: 3}
	circle1 := &Circle{Dot: Dot{x: 10, y: 10}, radius: 5}
	circle2 := &Circle{Dot: Dot{x: 15, y: 20}, radius: 8}

	all.Add(dot1)
	all.Add(dot2)

	subGroup := &CompoundGraphic{}
	subGroup.Add(circle1)
	subGroup.Add(circle2)

	all.Add(subGroup)

	fmt.Print(all.Draw())

	fmt.Println("\nMoving all graphics by (10, 10):")
	all.Move(10, 10)
	fmt.Print(all.Draw())
}
