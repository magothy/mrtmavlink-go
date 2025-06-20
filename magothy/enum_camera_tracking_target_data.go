//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"strings"
	"fmt"
)
// Camera tracking target data (shows where tracked target is within image)
type CAMERA_TRACKING_TARGET_DATA uint64

const (
	// No target data
	CAMERA_TRACKING_TARGET_NONE CAMERA_TRACKING_TARGET_DATA = 0
	// Target data embedded in image data (proprietary)
	CAMERA_TRACKING_TARGET_EMBEDDED CAMERA_TRACKING_TARGET_DATA = 1
	// Target data rendered in image
	CAMERA_TRACKING_TARGET_RENDERED CAMERA_TRACKING_TARGET_DATA = 2
	// Target data within status message (Point or Rectangle)
	CAMERA_TRACKING_TARGET_IN_STATUS CAMERA_TRACKING_TARGET_DATA = 4
)

var labels_CAMERA_TRACKING_TARGET_DATA = map[CAMERA_TRACKING_TARGET_DATA]string{
	CAMERA_TRACKING_TARGET_NONE: "CAMERA_TRACKING_TARGET_NONE",
	CAMERA_TRACKING_TARGET_EMBEDDED: "CAMERA_TRACKING_TARGET_EMBEDDED",
	CAMERA_TRACKING_TARGET_RENDERED: "CAMERA_TRACKING_TARGET_RENDERED",
	CAMERA_TRACKING_TARGET_IN_STATUS: "CAMERA_TRACKING_TARGET_IN_STATUS",
}

var values_CAMERA_TRACKING_TARGET_DATA = map[string]CAMERA_TRACKING_TARGET_DATA{
	"CAMERA_TRACKING_TARGET_NONE": CAMERA_TRACKING_TARGET_NONE,
	"CAMERA_TRACKING_TARGET_EMBEDDED": CAMERA_TRACKING_TARGET_EMBEDDED,
	"CAMERA_TRACKING_TARGET_RENDERED": CAMERA_TRACKING_TARGET_RENDERED,
	"CAMERA_TRACKING_TARGET_IN_STATUS": CAMERA_TRACKING_TARGET_IN_STATUS,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAMERA_TRACKING_TARGET_DATA) MarshalText() ([]byte, error) {
	if e == 0 {
		return []byte("0"), nil
	}
	var names []string
	for i := 0; i < 4; i++ {
		mask := CAMERA_TRACKING_TARGET_DATA(1 << i)
		if e&mask == mask {
			names = append(names, labels_CAMERA_TRACKING_TARGET_DATA[mask])
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAMERA_TRACKING_TARGET_DATA) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask CAMERA_TRACKING_TARGET_DATA
	for _, label := range labels {
		if value, ok := values_CAMERA_TRACKING_TARGET_DATA[label]; ok {
			mask |= value
		} else if value, err := strconv.Atoi(label); err == nil {
			mask |= CAMERA_TRACKING_TARGET_DATA(value)
		} else {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e CAMERA_TRACKING_TARGET_DATA) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
