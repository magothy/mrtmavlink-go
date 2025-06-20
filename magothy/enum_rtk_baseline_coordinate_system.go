//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"fmt"
)
// RTK GPS baseline coordinate system, used for RTK corrections
type RTK_BASELINE_COORDINATE_SYSTEM uint64

const (
	// Earth-centered, Earth-fixed
	RTK_BASELINE_COORDINATE_SYSTEM_ECEF RTK_BASELINE_COORDINATE_SYSTEM = 0
	// RTK basestation centered, north, east, down
	RTK_BASELINE_COORDINATE_SYSTEM_NED RTK_BASELINE_COORDINATE_SYSTEM = 1
)

var labels_RTK_BASELINE_COORDINATE_SYSTEM = map[RTK_BASELINE_COORDINATE_SYSTEM]string{
	RTK_BASELINE_COORDINATE_SYSTEM_ECEF: "RTK_BASELINE_COORDINATE_SYSTEM_ECEF",
	RTK_BASELINE_COORDINATE_SYSTEM_NED: "RTK_BASELINE_COORDINATE_SYSTEM_NED",
}

var values_RTK_BASELINE_COORDINATE_SYSTEM = map[string]RTK_BASELINE_COORDINATE_SYSTEM{
	"RTK_BASELINE_COORDINATE_SYSTEM_ECEF": RTK_BASELINE_COORDINATE_SYSTEM_ECEF,
	"RTK_BASELINE_COORDINATE_SYSTEM_NED": RTK_BASELINE_COORDINATE_SYSTEM_NED,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e RTK_BASELINE_COORDINATE_SYSTEM) MarshalText() ([]byte, error) {
	if name, ok := labels_RTK_BASELINE_COORDINATE_SYSTEM[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *RTK_BASELINE_COORDINATE_SYSTEM) UnmarshalText(text []byte) error {
	if value, ok := values_RTK_BASELINE_COORDINATE_SYSTEM[string(text)]; ok {
	   *e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
	   *e = RTK_BASELINE_COORDINATE_SYSTEM(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e RTK_BASELINE_COORDINATE_SYSTEM) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
