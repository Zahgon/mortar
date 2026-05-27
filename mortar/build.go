package mortar

import (
	"os"
	"time"
)

// Some of these variables should be populated during build using LDFLAGS
var (
	// should be injected with LDFLAG
	gitCommit string
	// should be injected with LDFLAG
	version string
	// should be injected with LDFLAG
	buildTimestamp string // "2006-01-02T15:04:05Z07:00" defined in RFC3339
	// should be injected with LDFLAG
	buildTag string

	// initialized During init()
	initTime time.Time
	// initialized During init()
	hostname string
)

func init() {
	initTime = time.Now()
	if host, err := os.Hostname(); err == nil {
		hostname = host
	} else {
		hostname = err.Error()
	}
}

// Information is a struct that will hold all the statically injected information during build
type Information struct {
	GitCommit string       `json:"git_commit,omitempty"`
	Version   string       `json:"version,omitempty"`
	BuildTag  string       `json:"build_tag,omitempty"`
	BuildTime time.Time    `json:"build_time,omitempty"`
	InitTime  time.Time    `json:"init_time,omitempty"`
	UpTime    JSONDuration `json:"up_time,omitempty"`
	Hostname  string       `json:"hostname,omitempty"`
}

// GetBuildInformation returns this service build information
func GetBuildInformation(includeExplanations ...bool) (info Information) {
	_ = "STUB: not implemented"
	return *new(Information)
}

// Zero

// try to parse

// JSONDuration is an alias to time.Duration for Json marshaling
type JSONDuration time.Duration

// MarshalJSON for JsonDuration is a helper function to better marshal time.Duration
func (jd JSONDuration) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
