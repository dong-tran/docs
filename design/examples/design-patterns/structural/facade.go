package structural

import "fmt"

// Facade Pattern
// Provides a simplified interface to a complex subsystem.

// Complex subsystem classes
type CPU struct{}

func (c *CPU) Freeze() { fmt.Println("CPU: Freezing") }
func (c *CPU) Jump(position int) { fmt.Printf("CPU: Jumping to position %d\n", position) }
func (c *CPU) Execute() { fmt.Println("CPU: Executing") }

type Memory struct{}

func (m *Memory) Load(position int, data string) {
	fmt.Printf("Memory: Loading '%s' at position %d\n", data, position)
}

type HardDrive struct{}

func (h *HardDrive) Read(sector int, size int) string {
	fmt.Printf("HardDrive: Reading %d bytes from sector %d\n", size, sector)
	return "boot_data"
}

// Facade
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (c *ComputerFacade) Start() {
	fmt.Println("Computer: Starting up...")
	c.cpu.Freeze()
	c.memory.Load(0, c.hardDrive.Read(0, 1024))
	c.cpu.Jump(0)
	c.cpu.Execute()
	fmt.Println("Computer: Ready!")
}

// Real-world example: Video conversion
type VideoFile struct{ name string }
type OggCompressionCodec struct{}
type MPEG4CompressionCodec struct{}
type CodecFactory struct{}
type BitrateReader struct{}
type AudioMixer struct{}

func (c *CodecFactory) Extract(file *VideoFile) string {
	fmt.Printf("CodecFactory: Extracting codec from %s\n", file.name)
	return "codec"
}

func (b *BitrateReader) Read(file *VideoFile, codec string) string {
	fmt.Printf("BitrateReader: Reading bitrate for %s\n", file.name)
	return "bitrate"
}

func (b *BitrateReader) Convert(buffer, codec string) string {
	fmt.Printf("BitrateReader: Converting buffer to %s\n", codec)
	return "converted_buffer"
}

func (a *AudioMixer) Fix(result string) string {
	fmt.Println("AudioMixer: Fixing audio")
	return result + "_fixed"
}

// Facade for video conversion
type VideoConversionFacade struct{}

func (v *VideoConversionFacade) ConvertVideo(fileName, format string) string {
	fmt.Printf("\n=== Converting %s to %s ===\n", fileName, format)
	
	file := &VideoFile{name: fileName}
	sourceCodec := (&CodecFactory{}).Extract(file)
	
	var destinationCodec string
	if format == "mp4" {
		destinationCodec = "MPEG4"
	} else {
		destinationCodec = "OGG"
	}
	
	buffer := (&BitrateReader{}).Read(file, sourceCodec)
	result := (&BitrateReader{}).Convert(buffer, destinationCodec)
	result = (&AudioMixer{}).Fix(result)
	
	fmt.Println("=== Conversion complete ===\n")
	return result
}

func DemoFacade() {
	fmt.Println("=== Facade Pattern Demo ===\n")

	fmt.Println("1. Computer Startup:")
	computer := NewComputerFacade()
	computer.Start()

	fmt.Println("\n2. Video Conversion:")
	converter := &VideoConversionFacade{}
	converter.ConvertVideo("video.avi", "mp4")
	converter.ConvertVideo("another.mkv", "ogg")
}
