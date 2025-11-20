package behavioral

import "fmt"

// Command Pattern
// Turns a request into a stand-alone object containing all information about the request.

type Command interface {
	Execute()
	Undo()
}

// Receiver
type Light struct {
	isOn bool
}

func (l *Light) On() {
	l.isOn = true
	fmt.Println("Light is ON")
}

func (l *Light) Off() {
	l.isOn = false
	fmt.Println("Light is OFF")
}

// Concrete Commands
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.Off()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}

// Invoker
type RemoteControl struct {
	command Command
	history []Command
}

func (r *RemoteControl) SetCommand(cmd Command) {
	r.command = cmd
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
	r.history = append(r.history, r.command)
}

func (r *RemoteControl) PressUndo() {
	if len(r.history) > 0 {
		cmd := r.history[len(r.history)-1]
		cmd.Undo()
		r.history = r.history[:len(r.history)-1]
	}
}

// Real-world example: Text Editor
type TextEditor struct {
	text string
}

func (e *TextEditor) GetText() string {
	return e.text
}

func (e *TextEditor) Write(text string) {
	e.text += text
}

func (e *TextEditor) Delete(length int) {
	if length > len(e.text) {
		length = len(e.text)
	}
	e.text = e.text[:len(e.text)-length]
}

type WriteCommand struct {
	editor *TextEditor
	text   string
}

func (c *WriteCommand) Execute() {
	c.editor.Write(c.text)
	fmt.Printf("Wrote: '%s' -> Text: '%s'\n", c.text, c.editor.GetText())
}

func (c *WriteCommand) Undo() {
	c.editor.Delete(len(c.text))
	fmt.Printf("Undid write -> Text: '%s'\n", c.editor.GetText())
}

func DemoCommand() {
	fmt.Println("=== Command Pattern Demo ===\n")

	fmt.Println("1. Light Control:")
	light := &Light{}
	remote := &RemoteControl{}

	remote.SetCommand(&LightOnCommand{light: light})
	remote.PressButton()

	remote.SetCommand(&LightOffCommand{light: light})
	remote.PressButton()

	fmt.Println("\nUndo last command:")
	remote.PressUndo()

	fmt.Println("\n2. Text Editor:")
	editor := &TextEditor{}
	history := []Command{}

	cmd1 := &WriteCommand{editor: editor, text: "Hello "}
	cmd1.Execute()
	history = append(history, cmd1)

	cmd2 := &WriteCommand{editor: editor, text: "World!"}
	cmd2.Execute()
	history = append(history, cmd2)

	fmt.Println("\nUndoing commands:")
	for i := len(history) - 1; i >= 0; i-- {
		history[i].Undo()
	}
}
