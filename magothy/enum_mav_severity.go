//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"fmt"
)
// Indicates the severity level, generally used for status messages to indicate their relative urgency. Based on RFC-5424 using expanded definitions at: http://www.kiwisyslog.com/kb/info:-syslog-message-levels/.
type MAV_SEVERITY uint64

const (
	// System is unusable. This is a "panic" condition.
	MAV_SEVERITY_EMERGENCY MAV_SEVERITY = 0
	// Action should be taken immediately. Indicates error in non-critical systems.
	MAV_SEVERITY_ALERT MAV_SEVERITY = 1
	// Action must be taken immediately. Indicates failure in a primary system.
	MAV_SEVERITY_CRITICAL MAV_SEVERITY = 2
	// Indicates an error in secondary/redundant systems.
	MAV_SEVERITY_ERROR MAV_SEVERITY = 3
	// Indicates about a possible future error if this is not resolved within a given timeframe. Example would be a low battery warning.
	MAV_SEVERITY_WARNING MAV_SEVERITY = 4
	// An unusual event has occurred, though not an error condition. This should be investigated for the root cause.
	MAV_SEVERITY_NOTICE MAV_SEVERITY = 5
	// Normal operational messages. Useful for logging. No action is required for these messages.
	MAV_SEVERITY_INFO MAV_SEVERITY = 6
	// Useful non-operational messages that can assist in debugging. These should not occur during normal operation.
	MAV_SEVERITY_DEBUG MAV_SEVERITY = 7
)

var labels_MAV_SEVERITY = map[MAV_SEVERITY]string{
	MAV_SEVERITY_EMERGENCY: "MAV_SEVERITY_EMERGENCY",
	MAV_SEVERITY_ALERT: "MAV_SEVERITY_ALERT",
	MAV_SEVERITY_CRITICAL: "MAV_SEVERITY_CRITICAL",
	MAV_SEVERITY_ERROR: "MAV_SEVERITY_ERROR",
	MAV_SEVERITY_WARNING: "MAV_SEVERITY_WARNING",
	MAV_SEVERITY_NOTICE: "MAV_SEVERITY_NOTICE",
	MAV_SEVERITY_INFO: "MAV_SEVERITY_INFO",
	MAV_SEVERITY_DEBUG: "MAV_SEVERITY_DEBUG",
}

var values_MAV_SEVERITY = map[string]MAV_SEVERITY{
	"MAV_SEVERITY_EMERGENCY": MAV_SEVERITY_EMERGENCY,
	"MAV_SEVERITY_ALERT": MAV_SEVERITY_ALERT,
	"MAV_SEVERITY_CRITICAL": MAV_SEVERITY_CRITICAL,
	"MAV_SEVERITY_ERROR": MAV_SEVERITY_ERROR,
	"MAV_SEVERITY_WARNING": MAV_SEVERITY_WARNING,
	"MAV_SEVERITY_NOTICE": MAV_SEVERITY_NOTICE,
	"MAV_SEVERITY_INFO": MAV_SEVERITY_INFO,
	"MAV_SEVERITY_DEBUG": MAV_SEVERITY_DEBUG,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_SEVERITY) MarshalText() ([]byte, error) {
	if name, ok := labels_MAV_SEVERITY[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_SEVERITY) UnmarshalText(text []byte) error {
	if value, ok := values_MAV_SEVERITY[string(text)]; ok {
	   *e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
	   *e = MAV_SEVERITY(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_SEVERITY) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
