package behavioral

import "fmt"

// Memento Pattern - Saves and restores the previous state of an object.

type Memento struct {
	state string
}

type Editor struct {
	content string
}

func (e *Editor) Type(words string) {
	e.content += words
}

func (e *Editor) GetContent() string {
	return e.content
}

func (e *Editor) Save() *Memento {
	return &Memento{state: e.content}
}

func (e *Editor) Restore(m *Memento) {
	e.content = m.state
}

type History struct {
	mementos []*Memento
}

func (h *History) Push(m *Memento) {
	h.mementos = append(h.mementos, m)
}

func (h *History) Pop() *Memento {
	if len(h.mementos) == 0 {
		return nil
	}
	lastIndex := len(h.mementos) - 1
	memento := h.mementos[lastIndex]
	h.mementos = h.mementos[:lastIndex]
	return memento
}

func DemoMemento() {
	fmt.Println("=== Memento Pattern Demo ===\n")
	editor := &Editor{}
	history := &History{}
	editor.Type("First sentence. ")
	history.Push(editor.Save())
	fmt.Printf("Saved: '%s'\n", editor.GetContent())
	editor.Type("Second sentence. ")
	history.Push(editor.Save())
	fmt.Printf("Saved: '%s'\n", editor.GetContent())
	editor.Type("Third sentence.")
	fmt.Printf("Current: '%s'\n", editor.GetContent())
	fmt.Println("\nUndo:")
	editor.Restore(history.Pop())
	fmt.Printf("After undo: '%s'\n", editor.GetContent())
	fmt.Println("\nUndo again:")
	editor.Restore(history.Pop())
	fmt.Printf("After undo: '%s'\n", editor.GetContent())
}
