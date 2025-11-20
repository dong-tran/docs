package structural

import "fmt"

// Bridge Pattern
// Decouples an abstraction from its implementation so they can vary independently.

// Implementation interface
type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	GetVolume() int
	SetVolume(percent int)
	GetChannel() int
	SetChannel(channel int)
}

// Concrete Implementations
type TV struct {
	on      bool
	volume  int
	channel int
}

func (t *TV) IsEnabled() bool        { return t.on }
func (t *TV) Enable()                { t.on = true; fmt.Println("TV: Turned ON") }
func (t *TV) Disable()               { t.on = false; fmt.Println("TV: Turned OFF") }
func (t *TV) GetVolume() int         { return t.volume }
func (t *TV) SetVolume(percent int)  { t.volume = percent; fmt.Printf("TV: Volume set to %d%%\n", percent) }
func (t *TV) GetChannel() int        { return t.channel }
func (t *TV) SetChannel(channel int) { t.channel = channel; fmt.Printf("TV: Channel set to %d\n", channel) }

type Radio struct {
	on      bool
	volume  int
	channel int
}

func (r *Radio) IsEnabled() bool        { return r.on }
func (r *Radio) Enable()                { r.on = true; fmt.Println("Radio: Turned ON") }
func (r *Radio) Disable()               { r.on = false; fmt.Println("Radio: Turned OFF") }
func (r *Radio) GetVolume() int         { return r.volume }
func (r *Radio) SetVolume(percent int)  { r.volume = percent; fmt.Printf("Radio: Volume set to %d%%\n", percent) }
func (r *Radio) GetChannel() int        { return r.channel }
func (r *Radio) SetChannel(channel int) { r.channel = channel; fmt.Printf("Radio: Station set to %d\n", channel) }

// Abstraction
type Remote struct {
	device Device
}

func NewRemote(device Device) *Remote {
	return &Remote{device: device}
}

func (r *Remote) TogglePower() {
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

func (r *Remote) VolumeDown() {
	vol := r.device.GetVolume()
	r.device.SetVolume(vol - 10)
}

func (r *Remote) VolumeUp() {
	vol := r.device.GetVolume()
	r.device.SetVolume(vol + 10)
}

func (r *Remote) ChannelDown() {
	ch := r.device.GetChannel()
	r.device.SetChannel(ch - 1)
}

func (r *Remote) ChannelUp() {
	ch := r.device.GetChannel()
	r.device.SetChannel(ch + 1)
}

// Refined Abstraction
type AdvancedRemote struct {
	*Remote
}

func NewAdvancedRemote(device Device) *AdvancedRemote {
	return &AdvancedRemote{Remote: NewRemote(device)}
}

func (a *AdvancedRemote) Mute() {
	fmt.Println("Advanced Remote: Muting")
	a.device.SetVolume(0)
}

func (a *AdvancedRemote) GoToChannel(channel int) {
	fmt.Printf("Advanced Remote: Going to channel %d\n", channel)
	a.device.SetChannel(channel)
}

func DemoBridge() {
	fmt.Println("=== Bridge Pattern Demo ===\n")

	tv := &TV{}
	remote := NewRemote(tv)

	fmt.Println("Testing basic remote with TV:")
	remote.TogglePower()
	remote.VolumeUp()
	remote.ChannelUp()

	fmt.Println("\nTesting advanced remote with TV:")
	advancedRemote := NewAdvancedRemote(tv)
	advancedRemote.Mute()
	advancedRemote.GoToChannel(42)

	fmt.Println("\nTesting with Radio:")
	radio := &Radio{}
	radioRemote := NewAdvancedRemote(radio)
	radioRemote.TogglePower()
	radioRemote.VolumeUp()
	radioRemote.GoToChannel(101)
}
