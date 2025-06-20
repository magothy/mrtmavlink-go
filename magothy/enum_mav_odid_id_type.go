//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"fmt"
)
type MAV_ODID_ID_TYPE uint64

const (
	// No type defined.
	MAV_ODID_ID_TYPE_NONE MAV_ODID_ID_TYPE = 0
	// Manufacturer Serial Number (ANSI/CTA-2063 format).
	MAV_ODID_ID_TYPE_SERIAL_NUMBER MAV_ODID_ID_TYPE = 1
	// CAA (Civil Aviation Authority) registered ID. Format: [ICAO Country Code].[CAA Assigned ID].
	MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID MAV_ODID_ID_TYPE = 2
	// UTM (Unmanned Traffic Management) assigned UUID (RFC4122).
	MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID MAV_ODID_ID_TYPE = 3
)

var labels_MAV_ODID_ID_TYPE = map[MAV_ODID_ID_TYPE]string{
	MAV_ODID_ID_TYPE_NONE: "MAV_ODID_ID_TYPE_NONE",
	MAV_ODID_ID_TYPE_SERIAL_NUMBER: "MAV_ODID_ID_TYPE_SERIAL_NUMBER",
	MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID: "MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID",
	MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID: "MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID",
}

var values_MAV_ODID_ID_TYPE = map[string]MAV_ODID_ID_TYPE{
	"MAV_ODID_ID_TYPE_NONE": MAV_ODID_ID_TYPE_NONE,
	"MAV_ODID_ID_TYPE_SERIAL_NUMBER": MAV_ODID_ID_TYPE_SERIAL_NUMBER,
	"MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID": MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID,
	"MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID": MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_ID_TYPE) MarshalText() ([]byte, error) {
	if name, ok := labels_MAV_ODID_ID_TYPE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_ID_TYPE) UnmarshalText(text []byte) error {
	if value, ok := values_MAV_ODID_ID_TYPE[string(text)]; ok {
	   *e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
	   *e = MAV_ODID_ID_TYPE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_ID_TYPE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
