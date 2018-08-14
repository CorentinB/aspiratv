package providers

import (
	"time"
)

// Show strcture hanlde show details.
// It is shared among all packages
type Show struct {
	ID           string        // Show ID
	Title        string        // Episode title
	Show         string        // Show name
	Pitch        string        // Pitch on the episode
	Season       string        // Season
	Episode      string        // Episode
	Channel      string        // Channel
	AirDate      time.Time     // Broadcasting date
	Duration     time.Duration // Duration of the show
	ShowURL      string        // url to the show page at provider
	StreamURL    string        // url of the video stream
	ThumbnailURL string        // direct url to the thumbnal of the show provided by the provider
	Detailed     bool          // False means some details can be requested
	DRM          bool          // True when video is protected with DRM
	Category     string        // Show's category
	Provider     string        // provider's name
}