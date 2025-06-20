//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"fmt"
)
// Enumeration of battery functions
type MAV_BATTERY_FUNCTION uint64

const (
	// Battery function is unknown
	MAV_BATTERY_FUNCTION_UNKNOWN MAV_BATTERY_FUNCTION = 0
	// Battery supports all flight systems
	MAV_BATTERY_FUNCTION_ALL MAV_BATTERY_FUNCTION = 1
	// Battery for the propulsion system
	MAV_BATTERY_FUNCTION_PROPULSION MAV_BATTERY_FUNCTION = 2
	// Avionics battery
	MAV_BATTERY_FUNCTION_AVIONICS MAV_BATTERY_FUNCTION = 3
	// Payload battery
	MAV_BATTERY_TYPE_PAYLOAD MAV_BATTERY_FUNCTION = 4
)

var labels_MAV_BATTERY_FUNCTION = map[MAV_BATTERY_FUNCTION]string{
	MAV_BATTERY_FUNCTION_UNKNOWN: "MAV_BATTERY_FUNCTION_UNKNOWN",
	MAV_BATTERY_FUNCTION_ALL: "MAV_BATTERY_FUNCTION_ALL",
	MAV_BATTERY_FUNCTION_PROPULSION: "MAV_BATTERY_FUNCTION_PROPULSION",
	MAV_BATTERY_FUNCTION_AVIONICS: "MAV_BATTERY_FUNCTION_AVIONICS",
	MAV_BATTERY_TYPE_PAYLOAD: "MAV_BATTERY_TYPE_PAYLOAD",
}

var values_MAV_BATTERY_FUNCTION = map[string]MAV_BATTERY_FUNCTION{
	"MAV_BATTERY_FUNCTION_UNKNOWN": MAV_BATTERY_FUNCTION_UNKNOWN,
	"MAV_BATTERY_FUNCTION_ALL": MAV_BATTERY_FUNCTION_ALL,
	"MAV_BATTERY_FUNCTION_PROPULSION": MAV_BATTERY_FUNCTION_PROPULSION,
	"MAV_BATTERY_FUNCTION_AVIONICS": MAV_BATTERY_FUNCTION_AVIONICS,
	"MAV_BATTERY_TYPE_PAYLOAD": MAV_BATTERY_TYPE_PAYLOAD,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_BATTERY_FUNCTION) MarshalText() ([]byte, error) {
	if name, ok := labels_MAV_BATTERY_FUNCTION[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_BATTERY_FUNCTION) UnmarshalText(text []byte) error {
	if value, ok := values_MAV_BATTERY_FUNCTION[string(text)]; ok {
	   *e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
	   *e = MAV_BATTERY_FUNCTION(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_BATTERY_FUNCTION) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
