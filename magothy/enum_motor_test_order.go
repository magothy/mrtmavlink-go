//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"fmt"
)
type MOTOR_TEST_ORDER uint64

const (
	// default autopilot motor test method
	MOTOR_TEST_ORDER_DEFAULT MOTOR_TEST_ORDER = 0
	// motor numbers are specified as their index in a predefined vehicle-specific sequence
	MOTOR_TEST_ORDER_SEQUENCE MOTOR_TEST_ORDER = 1
	// motor numbers are specified as the output as labeled on the board
	MOTOR_TEST_ORDER_BOARD MOTOR_TEST_ORDER = 2
)

var labels_MOTOR_TEST_ORDER = map[MOTOR_TEST_ORDER]string{
	MOTOR_TEST_ORDER_DEFAULT: "MOTOR_TEST_ORDER_DEFAULT",
	MOTOR_TEST_ORDER_SEQUENCE: "MOTOR_TEST_ORDER_SEQUENCE",
	MOTOR_TEST_ORDER_BOARD: "MOTOR_TEST_ORDER_BOARD",
}

var values_MOTOR_TEST_ORDER = map[string]MOTOR_TEST_ORDER{
	"MOTOR_TEST_ORDER_DEFAULT": MOTOR_TEST_ORDER_DEFAULT,
	"MOTOR_TEST_ORDER_SEQUENCE": MOTOR_TEST_ORDER_SEQUENCE,
	"MOTOR_TEST_ORDER_BOARD": MOTOR_TEST_ORDER_BOARD,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MOTOR_TEST_ORDER) MarshalText() ([]byte, error) {
	if name, ok := labels_MOTOR_TEST_ORDER[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MOTOR_TEST_ORDER) UnmarshalText(text []byte) error {
	if value, ok := values_MOTOR_TEST_ORDER[string(text)]; ok {
	   *e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
	   *e = MOTOR_TEST_ORDER(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e MOTOR_TEST_ORDER) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
