package creational

import "fmt"

// Abstract Factory Pattern
// Provides an interface for creating families of related or dependent objects
// without specifying their concrete classes.

// Abstract Products
type Button interface {
	Render() string
	OnClick()
}

type Checkbox interface {
	Render() string
	Toggle()
}

// Concrete Products - Windows Style
type WindowsButton struct {
	text string
}

func (b *WindowsButton) Render() string {
	return fmt.Sprintf("[Windows Button: %s]", b.text)
}

func (b *WindowsButton) OnClick() {
	fmt.Println("Windows button clicked with sound effect")
}

type WindowsCheckbox struct {
	label   string
	checked bool
}

func (c *WindowsCheckbox) Render() string {
	check := " "
	if c.checked {
		check = "X"
	}
	return fmt.Sprintf("[%s] %s", check, c.label)
}

func (c *WindowsCheckbox) Toggle() {
	c.checked = !c.checked
	fmt.Printf("Windows checkbox %s\n", map[bool]string{true: "checked", false: "unchecked"}[c.checked])
}

// Concrete Products - macOS Style
type MacButton struct {
	text string
}

func (b *MacButton) Render() string {
	return fmt.Sprintf("◉ %s ◉", b.text)
}

func (b *MacButton) OnClick() {
	fmt.Println("Mac button clicked with elegant animation")
}

type MacCheckbox struct {
	label   string
	checked bool
}

func (c *MacCheckbox) Render() string {
	check := "○"
	if c.checked {
		check = "●"
	}
	return fmt.Sprintf("%s %s", check, c.label)
}

func (c *MacCheckbox) Toggle() {
	c.checked = !c.checked
	fmt.Printf("Mac checkbox %s with smooth transition\n", map[bool]string{true: "checked", false: "unchecked"}[c.checked])
}

// Abstract Factory Interface
type GUIFactory interface {
	CreateButton(text string) Button
	CreateCheckbox(label string) Checkbox
}

// Concrete Factory - Windows
type WindowsFactory struct{}

func (f *WindowsFactory) CreateButton(text string) Button {
	return &WindowsButton{text: text}
}

func (f *WindowsFactory) CreateCheckbox(label string) Checkbox {
	return &WindowsCheckbox{label: label, checked: false}
}

// Concrete Factory - macOS
type MacFactory struct{}

func (f *MacFactory) CreateButton(text string) Button {
	return &MacButton{text: text}
}

func (f *MacFactory) CreateCheckbox(label string) Checkbox {
	return &MacCheckbox{label: label, checked: false}
}

// Application that uses the factory
type Application struct {
	factory GUIFactory
}

func NewApplication(factory GUIFactory) *Application {
	return &Application{factory: factory}
}

func (app *Application) RenderUI() {
	button := app.factory.CreateButton("Submit")
	checkbox := app.factory.CreateCheckbox("Accept terms")

	fmt.Println("Rendering UI:")
	fmt.Println(button.Render())
	fmt.Println(checkbox.Render())
	
	button.OnClick()
	checkbox.Toggle()
	fmt.Println(checkbox.Render())
}

// Example usage demonstrating the pattern
func DemoAbstractFactory() {
	fmt.Println("=== Abstract Factory Pattern Demo ===\n")

	// Create Windows application
	fmt.Println("Creating Windows Application:")
	windowsApp := NewApplication(&WindowsFactory{})
	windowsApp.RenderUI()

	fmt.Println("\n---\n")

	// Create Mac application
	fmt.Println("Creating Mac Application:")
	macApp := NewApplication(&MacFactory{})
	macApp.RenderUI()

	// Real-world example: Database connections with different drivers
	fmt.Println("\n\n=== Real-World Example: Database Drivers ===\n")

	// PostgreSQL family
	type Connection interface {
		Connect() string
		Query(sql string) string
	}

	type Transaction interface {
		Begin() string
		Commit() string
		Rollback() string
	}

	type DatabaseFactory interface {
		CreateConnection(host string) Connection
		CreateTransaction() Transaction
	}

	// PostgreSQL implementations
	type PostgresConnection struct{ host string }
	func (c *PostgresConnection) Connect() string { return fmt.Sprintf("Connected to PostgreSQL at %s", c.host) }
	func (c *PostgresConnection) Query(sql string) string { return fmt.Sprintf("PostgreSQL executing: %s", sql) }

	type PostgresTransaction struct{}
	func (t *PostgresTransaction) Begin() string { return "PostgreSQL: BEGIN" }
	func (t *PostgresTransaction) Commit() string { return "PostgreSQL: COMMIT" }
	func (t *PostgresTransaction) Rollback() string { return "PostgreSQL: ROLLBACK" }

	type PostgresFactory struct{}
	func (f *PostgresFactory) CreateConnection(host string) Connection { return &PostgresConnection{host} }
	func (f *PostgresFactory) CreateTransaction() Transaction { return &PostgresTransaction{} }

	// MySQL implementations
	type MySQLConnection struct{ host string }
	func (c *MySQLConnection) Connect() string { return fmt.Sprintf("Connected to MySQL at %s", c.host) }
	func (c *MySQLConnection) Query(sql string) string { return fmt.Sprintf("MySQL executing: %s", sql) }

	type MySQLTransaction struct{}
	func (t *MySQLTransaction) Begin() string { return "MySQL: START TRANSACTION" }
	func (t *MySQLTransaction) Commit() string { return "MySQL: COMMIT" }
	func (t *MySQLTransaction) Rollback() string { return "MySQL: ROLLBACK" }

	type MySQLFactory struct{}
	func (f *MySQLFactory) CreateConnection(host string) Connection { return &MySQLConnection{host} }
	func (f *MySQLFactory) CreateTransaction() Transaction { return &MySQLTransaction{} }

	// Use the factories
	var dbFactory DatabaseFactory

	dbFactory = &PostgresFactory{}
	conn := dbFactory.CreateConnection("localhost:5432")
	tx := dbFactory.CreateTransaction()
	fmt.Println(conn.Connect())
	fmt.Println(conn.Query("SELECT * FROM users"))
	fmt.Println(tx.Begin())
	fmt.Println(tx.Commit())

	fmt.Println()

	dbFactory = &MySQLFactory{}
	conn = dbFactory.CreateConnection("localhost:3306")
	tx = dbFactory.CreateTransaction()
	fmt.Println(conn.Connect())
	fmt.Println(conn.Query("SELECT * FROM users"))
	fmt.Println(tx.Begin())
	fmt.Println(tx.Commit())
}
