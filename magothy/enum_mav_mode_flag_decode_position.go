//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"strings"
	"fmt"
)
// These values encode the bit positions of the decode position. These values can be used to read the value of a flag bit by combining the base_mode variable with AND with the flag position value. The result will be either 0 or 1, depending on if the flag is set or not.
type MAV_MODE_FLAG_DECODE_POSITION uint64

const (
	// First bit:  10000000
	MAV_MODE_FLAG_DECODE_POSITION_SAFETY MAV_MODE_FLAG_DECODE_POSITION = 128
	// Second bit: 01000000
	MAV_MODE_FLAG_DECODE_POSITION_MANUAL MAV_MODE_FLAG_DECODE_POSITION = 64
	// Third bit:  00100000
	MAV_MODE_FLAG_DECODE_POSITION_HIL MAV_MODE_FLAG_DECODE_POSITION = 32
	// Fourth bit: 00010000
	MAV_MODE_FLAG_DECODE_POSITION_STABILIZE MAV_MODE_FLAG_DECODE_POSITION = 16
	// Fifth bit:  00001000
	MAV_MODE_FLAG_DECODE_POSITION_GUIDED MAV_MODE_FLAG_DECODE_POSITION = 8
	// Sixth bit:   00000100
	MAV_MODE_FLAG_DECODE_POSITION_AUTO MAV_MODE_FLAG_DECODE_POSITION = 4
	// Seventh bit: 00000010
	MAV_MODE_FLAG_DECODE_POSITION_TEST MAV_MODE_FLAG_DECODE_POSITION = 2
	// Eighth bit: 00000001
	MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE MAV_MODE_FLAG_DECODE_POSITION = 1
)

var labels_MAV_MODE_FLAG_DECODE_POSITION = map[MAV_MODE_FLAG_DECODE_POSITION]string{
	MAV_MODE_FLAG_DECODE_POSITION_SAFETY: "MAV_MODE_FLAG_DECODE_POSITION_SAFETY",
	MAV_MODE_FLAG_DECODE_POSITION_MANUAL: "MAV_MODE_FLAG_DECODE_POSITION_MANUAL",
	MAV_MODE_FLAG_DECODE_POSITION_HIL: "MAV_MODE_FLAG_DECODE_POSITION_HIL",
	MAV_MODE_FLAG_DECODE_POSITION_STABILIZE: "MAV_MODE_FLAG_DECODE_POSITION_STABILIZE",
	MAV_MODE_FLAG_DECODE_POSITION_GUIDED: "MAV_MODE_FLAG_DECODE_POSITION_GUIDED",
	MAV_MODE_FLAG_DECODE_POSITION_AUTO: "MAV_MODE_FLAG_DECODE_POSITION_AUTO",
	MAV_MODE_FLAG_DECODE_POSITION_TEST: "MAV_MODE_FLAG_DECODE_POSITION_TEST",
	MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE: "MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE",
}

var values_MAV_MODE_FLAG_DECODE_POSITION = map[string]MAV_MODE_FLAG_DECODE_POSITION{
	"MAV_MODE_FLAG_DECODE_POSITION_SAFETY": MAV_MODE_FLAG_DECODE_POSITION_SAFETY,
	"MAV_MODE_FLAG_DECODE_POSITION_MANUAL": MAV_MODE_FLAG_DECODE_POSITION_MANUAL,
	"MAV_MODE_FLAG_DECODE_POSITION_HIL": MAV_MODE_FLAG_DECODE_POSITION_HIL,
	"MAV_MODE_FLAG_DECODE_POSITION_STABILIZE": MAV_MODE_FLAG_DECODE_POSITION_STABILIZE,
	"MAV_MODE_FLAG_DECODE_POSITION_GUIDED": MAV_MODE_FLAG_DECODE_POSITION_GUIDED,
	"MAV_MODE_FLAG_DECODE_POSITION_AUTO": MAV_MODE_FLAG_DECODE_POSITION_AUTO,
	"MAV_MODE_FLAG_DECODE_POSITION_TEST": MAV_MODE_FLAG_DECODE_POSITION_TEST,
	"MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE": MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_MODE_FLAG_DECODE_POSITION) MarshalText() ([]byte, error) {
	if e == 0 {
		return []byte("0"), nil
	}
	var names []string
	for i := 0; i < 8; i++ {
		mask := MAV_MODE_FLAG_DECODE_POSITION(1 << i)
		if e&mask == mask {
			names = append(names, labels_MAV_MODE_FLAG_DECODE_POSITION[mask])
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_MODE_FLAG_DECODE_POSITION) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_MODE_FLAG_DECODE_POSITION
	for _, label := range labels {
		if value, ok := values_MAV_MODE_FLAG_DECODE_POSITION[label]; ok {
			mask |= value
		} else if value, err := strconv.Atoi(label); err == nil {
			mask |= MAV_MODE_FLAG_DECODE_POSITION(value)
		} else {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_MODE_FLAG_DECODE_POSITION) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
