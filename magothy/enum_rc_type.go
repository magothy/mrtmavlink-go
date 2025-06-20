//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"fmt"
)
// RC type
type RC_TYPE uint64

const (
	// Spektrum DSM2
	RC_TYPE_SPEKTRUM_DSM2 RC_TYPE = 0
	// Spektrum DSMX
	RC_TYPE_SPEKTRUM_DSMX RC_TYPE = 1
)

var labels_RC_TYPE = map[RC_TYPE]string{
	RC_TYPE_SPEKTRUM_DSM2: "RC_TYPE_SPEKTRUM_DSM2",
	RC_TYPE_SPEKTRUM_DSMX: "RC_TYPE_SPEKTRUM_DSMX",
}

var values_RC_TYPE = map[string]RC_TYPE{
	"RC_TYPE_SPEKTRUM_DSM2": RC_TYPE_SPEKTRUM_DSM2,
	"RC_TYPE_SPEKTRUM_DSMX": RC_TYPE_SPEKTRUM_DSMX,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e RC_TYPE) MarshalText() ([]byte, error) {
	if name, ok := labels_RC_TYPE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *RC_TYPE) UnmarshalText(text []byte) error {
	if value, ok := values_RC_TYPE[string(text)]; ok {
	   *e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
	   *e = RC_TYPE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e RC_TYPE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
