//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"fmt"
)
// WiFi Mode.
type WIFI_CONFIG_AP_MODE uint64

const (
	// WiFi mode is undefined.
	WIFI_CONFIG_AP_MODE_UNDEFINED WIFI_CONFIG_AP_MODE = 0
	// WiFi configured as an access point.
	WIFI_CONFIG_AP_MODE_AP WIFI_CONFIG_AP_MODE = 1
	// WiFi configured as a station connected to an existing local WiFi network.
	WIFI_CONFIG_AP_MODE_STATION WIFI_CONFIG_AP_MODE = 2
	// WiFi disabled.
	WIFI_CONFIG_AP_MODE_DISABLED WIFI_CONFIG_AP_MODE = 3
)

var labels_WIFI_CONFIG_AP_MODE = map[WIFI_CONFIG_AP_MODE]string{
	WIFI_CONFIG_AP_MODE_UNDEFINED: "WIFI_CONFIG_AP_MODE_UNDEFINED",
	WIFI_CONFIG_AP_MODE_AP: "WIFI_CONFIG_AP_MODE_AP",
	WIFI_CONFIG_AP_MODE_STATION: "WIFI_CONFIG_AP_MODE_STATION",
	WIFI_CONFIG_AP_MODE_DISABLED: "WIFI_CONFIG_AP_MODE_DISABLED",
}

var values_WIFI_CONFIG_AP_MODE = map[string]WIFI_CONFIG_AP_MODE{
	"WIFI_CONFIG_AP_MODE_UNDEFINED": WIFI_CONFIG_AP_MODE_UNDEFINED,
	"WIFI_CONFIG_AP_MODE_AP": WIFI_CONFIG_AP_MODE_AP,
	"WIFI_CONFIG_AP_MODE_STATION": WIFI_CONFIG_AP_MODE_STATION,
	"WIFI_CONFIG_AP_MODE_DISABLED": WIFI_CONFIG_AP_MODE_DISABLED,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e WIFI_CONFIG_AP_MODE) MarshalText() ([]byte, error) {
	if name, ok := labels_WIFI_CONFIG_AP_MODE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *WIFI_CONFIG_AP_MODE) UnmarshalText(text []byte) error {
	if value, ok := values_WIFI_CONFIG_AP_MODE[string(text)]; ok {
	   *e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
	   *e = WIFI_CONFIG_AP_MODE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e WIFI_CONFIG_AP_MODE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
