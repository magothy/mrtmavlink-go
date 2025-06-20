//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"fmt"
)
// These flags encode the cellular network status
type CELLULAR_STATUS_FLAG uint64

const (
	// State unknown or not reportable.
	CELLULAR_STATUS_FLAG_UNKNOWN CELLULAR_STATUS_FLAG = 0
	// Modem is unusable
	CELLULAR_STATUS_FLAG_FAILED CELLULAR_STATUS_FLAG = 1
	// Modem is being initialized
	CELLULAR_STATUS_FLAG_INITIALIZING CELLULAR_STATUS_FLAG = 2
	// Modem is locked
	CELLULAR_STATUS_FLAG_LOCKED CELLULAR_STATUS_FLAG = 3
	// Modem is not enabled and is powered down
	CELLULAR_STATUS_FLAG_DISABLED CELLULAR_STATUS_FLAG = 4
	// Modem is currently transitioning to the CELLULAR_STATUS_FLAG_DISABLED state
	CELLULAR_STATUS_FLAG_DISABLING CELLULAR_STATUS_FLAG = 5
	// Modem is currently transitioning to the CELLULAR_STATUS_FLAG_ENABLED state
	CELLULAR_STATUS_FLAG_ENABLING CELLULAR_STATUS_FLAG = 6
	// Modem is enabled and powered on but not registered with a network provider and not available for data connections
	CELLULAR_STATUS_FLAG_ENABLED CELLULAR_STATUS_FLAG = 7
	// Modem is searching for a network provider to register
	CELLULAR_STATUS_FLAG_SEARCHING CELLULAR_STATUS_FLAG = 8
	// Modem is registered with a network provider, and data connections and messaging may be available for use
	CELLULAR_STATUS_FLAG_REGISTERED CELLULAR_STATUS_FLAG = 9
	// Modem is disconnecting and deactivating the last active packet data bearer. This state will not be entered if more than one packet data bearer is active and one of the active bearers is deactivated
	CELLULAR_STATUS_FLAG_DISCONNECTING CELLULAR_STATUS_FLAG = 10
	// Modem is activating and connecting the first packet data bearer. Subsequent bearer activations when another bearer is already active do not cause this state to be entered
	CELLULAR_STATUS_FLAG_CONNECTING CELLULAR_STATUS_FLAG = 11
	// One or more packet data bearers is active and connected
	CELLULAR_STATUS_FLAG_CONNECTED CELLULAR_STATUS_FLAG = 12
)

var labels_CELLULAR_STATUS_FLAG = map[CELLULAR_STATUS_FLAG]string{
	CELLULAR_STATUS_FLAG_UNKNOWN: "CELLULAR_STATUS_FLAG_UNKNOWN",
	CELLULAR_STATUS_FLAG_FAILED: "CELLULAR_STATUS_FLAG_FAILED",
	CELLULAR_STATUS_FLAG_INITIALIZING: "CELLULAR_STATUS_FLAG_INITIALIZING",
	CELLULAR_STATUS_FLAG_LOCKED: "CELLULAR_STATUS_FLAG_LOCKED",
	CELLULAR_STATUS_FLAG_DISABLED: "CELLULAR_STATUS_FLAG_DISABLED",
	CELLULAR_STATUS_FLAG_DISABLING: "CELLULAR_STATUS_FLAG_DISABLING",
	CELLULAR_STATUS_FLAG_ENABLING: "CELLULAR_STATUS_FLAG_ENABLING",
	CELLULAR_STATUS_FLAG_ENABLED: "CELLULAR_STATUS_FLAG_ENABLED",
	CELLULAR_STATUS_FLAG_SEARCHING: "CELLULAR_STATUS_FLAG_SEARCHING",
	CELLULAR_STATUS_FLAG_REGISTERED: "CELLULAR_STATUS_FLAG_REGISTERED",
	CELLULAR_STATUS_FLAG_DISCONNECTING: "CELLULAR_STATUS_FLAG_DISCONNECTING",
	CELLULAR_STATUS_FLAG_CONNECTING: "CELLULAR_STATUS_FLAG_CONNECTING",
	CELLULAR_STATUS_FLAG_CONNECTED: "CELLULAR_STATUS_FLAG_CONNECTED",
}

var values_CELLULAR_STATUS_FLAG = map[string]CELLULAR_STATUS_FLAG{
	"CELLULAR_STATUS_FLAG_UNKNOWN": CELLULAR_STATUS_FLAG_UNKNOWN,
	"CELLULAR_STATUS_FLAG_FAILED": CELLULAR_STATUS_FLAG_FAILED,
	"CELLULAR_STATUS_FLAG_INITIALIZING": CELLULAR_STATUS_FLAG_INITIALIZING,
	"CELLULAR_STATUS_FLAG_LOCKED": CELLULAR_STATUS_FLAG_LOCKED,
	"CELLULAR_STATUS_FLAG_DISABLED": CELLULAR_STATUS_FLAG_DISABLED,
	"CELLULAR_STATUS_FLAG_DISABLING": CELLULAR_STATUS_FLAG_DISABLING,
	"CELLULAR_STATUS_FLAG_ENABLING": CELLULAR_STATUS_FLAG_ENABLING,
	"CELLULAR_STATUS_FLAG_ENABLED": CELLULAR_STATUS_FLAG_ENABLED,
	"CELLULAR_STATUS_FLAG_SEARCHING": CELLULAR_STATUS_FLAG_SEARCHING,
	"CELLULAR_STATUS_FLAG_REGISTERED": CELLULAR_STATUS_FLAG_REGISTERED,
	"CELLULAR_STATUS_FLAG_DISCONNECTING": CELLULAR_STATUS_FLAG_DISCONNECTING,
	"CELLULAR_STATUS_FLAG_CONNECTING": CELLULAR_STATUS_FLAG_CONNECTING,
	"CELLULAR_STATUS_FLAG_CONNECTED": CELLULAR_STATUS_FLAG_CONNECTED,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CELLULAR_STATUS_FLAG) MarshalText() ([]byte, error) {
	if name, ok := labels_CELLULAR_STATUS_FLAG[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CELLULAR_STATUS_FLAG) UnmarshalText(text []byte) error {
	if value, ok := values_CELLULAR_STATUS_FLAG[string(text)]; ok {
	   *e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
	   *e = CELLULAR_STATUS_FLAG(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e CELLULAR_STATUS_FLAG) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
