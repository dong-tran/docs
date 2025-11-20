package structural

// Adapter - Structural Pattern
// Allows incompatible interfaces to work together

// Target interface
type MediaPlayer interface {
	Play(filename string)
}

// Adaptee - incompatible interface
type AdvancedMediaPlayer interface {
	PlayVLC(filename string)
	PlayMP4(filename string)
}

type VLCPlayer struct{}

func (v *VLCPlayer) PlayVLC(filename string) {
	// Play VLC
}

func (v *VLCPlayer) PlayMP4(filename string) {}

type MP4Player struct{}

func (m *MP4Player) PlayVLC(filename string) {}

func (m *MP4Player) PlayMP4(filename string) {
	// Play MP4
}

// Adapter
type MediaAdapter struct {
	advancedPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	if audioType == "vlc" {
		return &MediaAdapter{advancedPlayer: &VLCPlayer{}}
	} else if audioType == "mp4" {
		return &MediaAdapter{advancedPlayer: &MP4Player{}}
	}
	return nil
}

func (a *MediaAdapter) Play(filename string) {
	// Adapt the interface
	if a.advancedPlayer != nil {
		a.advancedPlayer.PlayMP4(filename)
	}
}
