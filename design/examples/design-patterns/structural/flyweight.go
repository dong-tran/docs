package structural

import "fmt"

// Flyweight Pattern
// Uses sharing to support large numbers of fine-grained objects efficiently.
// Separates intrinsic state (shared) from extrinsic state (unique).

// Flyweight interface
type TreeType struct {
	name    string // intrinsic
	color   string // intrinsic
	texture string // intrinsic
}

func (t *TreeType) Draw(x, y int) {
	fmt.Printf("Drawing %s tree at (%d, %d) with %s color\n", t.name, x, y, t.color)
}

// Flyweight factory
type TreeFactory struct {
	treeTypes map[string]*TreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{
		treeTypes: make(map[string]*TreeType),
	}
}

func (f *TreeFactory) GetTreeType(name, color, texture string) *TreeType {
	key := name + "_" + color + "_" + texture
	
	if treeType, exists := f.treeTypes[key]; exists {
		return treeType
	}
	
	fmt.Printf("Creating new TreeType: %s\n", key)
	treeType := &TreeType{name: name, color: color, texture: texture}
	f.treeTypes[key] = treeType
	return treeType
}

func (f *TreeFactory) GetTotalTypes() int {
	return len(f.treeTypes)
}

// Context class that uses flyweight
type Tree struct {
	x        int       // extrinsic
	y        int       // extrinsic
	treeType *TreeType // flyweight
}

func (t *Tree) Draw() {
	t.treeType.Draw(t.x, t.y)
}

// Forest contains many trees
type Forest struct {
	trees       []*Tree
	treeFactory *TreeFactory
}

func NewForest() *Forest {
	return &Forest{
		trees:       make([]*Tree, 0),
		treeFactory: NewTreeFactory(),
	}
}

func (f *Forest) PlantTree(x, y int, name, color, texture string) {
	treeType := f.treeFactory.GetTreeType(name, color, texture)
	tree := &Tree{x: x, y: y, treeType: treeType}
	f.trees = append(f.trees, tree)
}

func (f *Forest) Draw() {
	for _, tree := range f.trees {
		tree.Draw()
	}
}

func (f *Forest) GetStats() {
	fmt.Printf("Forest has %d trees\n", len(f.trees))
	fmt.Printf("Forest uses only %d tree types (flyweights)\n", f.treeFactory.GetTotalTypes())
	fmt.Printf("Memory saved: %d tree objects share %d flyweights\n", 
len(f.trees), f.treeFactory.GetTotalTypes())
}

// Real-world example: Character formatting in text editor
type CharacterStyle struct {
	font   string
	size   int
	color  string
	bold   bool
	italic bool
}

type StyleFactory struct {
	styles map[string]*CharacterStyle
}

func NewStyleFactory() *StyleFactory {
	return &StyleFactory{styles: make(map[string]*CharacterStyle)}
}

func (sf *StyleFactory) GetStyle(font string, size int, color string, bold, italic bool) *CharacterStyle {
	key := fmt.Sprintf("%s_%d_%s_%v_%v", font, size, color, bold, italic)
	
	if style, exists := sf.styles[key]; exists {
		return style
	}
	
	style := &CharacterStyle{font: font, size: size, color: color, bold: bold, italic: italic}
	sf.styles[key] = style
	return style
}

type Character struct {
	char  rune
	style *CharacterStyle
}

func DemoFlyweight() {
	fmt.Println("=== Flyweight Pattern Demo ===\n")

	fmt.Println("1. Forest with many trees:")
	forest := NewForest()

	// Plant many trees, but use only a few types
	forest.PlantTree(1, 1, "Oak", "Green", "Rough")
	forest.PlantTree(2, 2, "Oak", "Green", "Rough")
	forest.PlantTree(3, 3, "Pine", "DarkGreen", "Smooth")
	forest.PlantTree(4, 4, "Oak", "Green", "Rough")
	forest.PlantTree(5, 5, "Pine", "DarkGreen", "Smooth")
	forest.PlantTree(6, 6, "Birch", "White", "Smooth")
	forest.PlantTree(7, 7, "Oak", "Green", "Rough")
	forest.PlantTree(8, 8, "Pine", "DarkGreen", "Smooth")

	fmt.Println("\nDrawing forest:")
	forest.Draw()

	fmt.Println()
	forest.GetStats()

	// Text editor example
	fmt.Println("\n2. Text Editor Character Formatting:")
	styleFactory := NewStyleFactory()

	text := "Hello World!"
	characters := make([]*Character, 0)

	// Most characters share same style
	defaultStyle := styleFactory.GetStyle("Arial", 12, "Black", false, false)
	for _, char := range text {
		characters = append(characters, &Character{char: char, style: defaultStyle})
	}

	// Some characters have different styles
	characters[0].style = styleFactory.GetStyle("Arial", 12, "Red", true, false)
	characters[6].style = styleFactory.GetStyle("Arial", 14, "Blue", false, true)

	fmt.Printf("Text has %d characters but uses only %d styles\n", 
len(characters), len(styleFactory.styles))
}
