package behavioral

import "fmt"

// Template Method Pattern - Defines skeleton of algorithm, deferring some steps to subclasses.

type DataProcessor interface {
	ReadData() string
	ProcessData(data string) string
	WriteData(data string)
	Process()
}

type BaseProcessor struct {
	processor DataProcessor
}

func (b *BaseProcessor) Process() {
	data := b.processor.ReadData()
	processed := b.processor.ProcessData(data)
	b.processor.WriteData(processed)
}

type CSVProcessor struct {
	BaseProcessor
	filename string
}

func NewCSVProcessor(filename string) *CSVProcessor {
	p := &CSVProcessor{filename: filename}
	p.BaseProcessor.processor = p
	return p
}

func (c *CSVProcessor) ReadData() string {
	fmt.Printf("Reading CSV file: %s\n", c.filename)
	return "csv_data"
}

func (c *CSVProcessor) ProcessData(data string) string {
	fmt.Printf("Processing CSV data: %s\n", data)
	return "processed_" + data
}

func (c *CSVProcessor) WriteData(data string) {
	fmt.Printf("Writing CSV result: %s\n", data)
}

func (c *CSVProcessor) Process() {
	c.BaseProcessor.Process()
}

type JSONProcessor struct {
	BaseProcessor
	filename string
}

func NewJSONProcessor(filename string) *JSONProcessor {
	p := &JSONProcessor{filename: filename}
	p.BaseProcessor.processor = p
	return p
}

func (j *JSONProcessor) ReadData() string {
	fmt.Printf("Reading JSON file: %s\n", j.filename)
	return "json_data"
}

func (j *JSONProcessor) ProcessData(data string) string {
	fmt.Printf("Processing JSON data: %s\n", data)
	return "processed_" + data
}

func (j *JSONProcessor) WriteData(data string) {
	fmt.Printf("Writing JSON result: %s\n", data)
}

func (j *JSONProcessor) Process() {
	j.BaseProcessor.Process()
}

func DemoTemplateMethod() {
	fmt.Println("=== Template Method Pattern Demo ===\n")
	fmt.Println("Processing CSV:")
	csvProcessor := NewCSVProcessor("data.csv")
	csvProcessor.Process()
	fmt.Println("\nProcessing JSON:")
	jsonProcessor := NewJSONProcessor("data.json")
	jsonProcessor.Process()
}
