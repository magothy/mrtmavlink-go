//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"strings"
	"fmt"
)
// Stream status flags (Bitmap)
type VIDEO_STREAM_STATUS_FLAGS uint64

const (
	// Stream is active (running)
	VIDEO_STREAM_STATUS_FLAGS_RUNNING VIDEO_STREAM_STATUS_FLAGS = 1
	// Stream is thermal imaging
	VIDEO_STREAM_STATUS_FLAGS_THERMAL VIDEO_STREAM_STATUS_FLAGS = 2
)

var labels_VIDEO_STREAM_STATUS_FLAGS = map[VIDEO_STREAM_STATUS_FLAGS]string{
	VIDEO_STREAM_STATUS_FLAGS_RUNNING: "VIDEO_STREAM_STATUS_FLAGS_RUNNING",
	VIDEO_STREAM_STATUS_FLAGS_THERMAL: "VIDEO_STREAM_STATUS_FLAGS_THERMAL",
}

var values_VIDEO_STREAM_STATUS_FLAGS = map[string]VIDEO_STREAM_STATUS_FLAGS{
	"VIDEO_STREAM_STATUS_FLAGS_RUNNING": VIDEO_STREAM_STATUS_FLAGS_RUNNING,
	"VIDEO_STREAM_STATUS_FLAGS_THERMAL": VIDEO_STREAM_STATUS_FLAGS_THERMAL,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e VIDEO_STREAM_STATUS_FLAGS) MarshalText() ([]byte, error) {
	if e == 0 {
		return []byte("0"), nil
	}
	var names []string
	for i := 0; i < 2; i++ {
		mask := VIDEO_STREAM_STATUS_FLAGS(1 << i)
		if e&mask == mask {
			names = append(names, labels_VIDEO_STREAM_STATUS_FLAGS[mask])
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *VIDEO_STREAM_STATUS_FLAGS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask VIDEO_STREAM_STATUS_FLAGS
	for _, label := range labels {
		if value, ok := values_VIDEO_STREAM_STATUS_FLAGS[label]; ok {
			mask |= value
		} else if value, err := strconv.Atoi(label); err == nil {
			mask |= VIDEO_STREAM_STATUS_FLAGS(value)
		} else {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e VIDEO_STREAM_STATUS_FLAGS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
