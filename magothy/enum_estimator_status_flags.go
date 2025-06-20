//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"strings"
	"fmt"
)
// Flags in ESTIMATOR_STATUS message
type ESTIMATOR_STATUS_FLAGS uint64

const (
	// True if the attitude estimate is good
	ESTIMATOR_ATTITUDE ESTIMATOR_STATUS_FLAGS = 1
	// True if the horizontal velocity estimate is good
	ESTIMATOR_VELOCITY_HORIZ ESTIMATOR_STATUS_FLAGS = 2
	// True if the  vertical velocity estimate is good
	ESTIMATOR_VELOCITY_VERT ESTIMATOR_STATUS_FLAGS = 4
	// True if the horizontal position (relative) estimate is good
	ESTIMATOR_POS_HORIZ_REL ESTIMATOR_STATUS_FLAGS = 8
	// True if the horizontal position (absolute) estimate is good
	ESTIMATOR_POS_HORIZ_ABS ESTIMATOR_STATUS_FLAGS = 16
	// True if the vertical position (absolute) estimate is good
	ESTIMATOR_POS_VERT_ABS ESTIMATOR_STATUS_FLAGS = 32
	// True if the vertical position (above ground) estimate is good
	ESTIMATOR_POS_VERT_AGL ESTIMATOR_STATUS_FLAGS = 64
	// True if the EKF is in a constant position mode and is not using external measurements (eg GPS or optical flow)
	ESTIMATOR_CONST_POS_MODE ESTIMATOR_STATUS_FLAGS = 128
	// True if the EKF has sufficient data to enter a mode that will provide a (relative) position estimate
	ESTIMATOR_PRED_POS_HORIZ_REL ESTIMATOR_STATUS_FLAGS = 256
	// True if the EKF has sufficient data to enter a mode that will provide a (absolute) position estimate
	ESTIMATOR_PRED_POS_HORIZ_ABS ESTIMATOR_STATUS_FLAGS = 512
	// True if the EKF has detected a GPS glitch
	ESTIMATOR_GPS_GLITCH ESTIMATOR_STATUS_FLAGS = 1024
	// True if the EKF has detected bad accelerometer data
	ESTIMATOR_ACCEL_ERROR ESTIMATOR_STATUS_FLAGS = 2048
)

var labels_ESTIMATOR_STATUS_FLAGS = map[ESTIMATOR_STATUS_FLAGS]string{
	ESTIMATOR_ATTITUDE: "ESTIMATOR_ATTITUDE",
	ESTIMATOR_VELOCITY_HORIZ: "ESTIMATOR_VELOCITY_HORIZ",
	ESTIMATOR_VELOCITY_VERT: "ESTIMATOR_VELOCITY_VERT",
	ESTIMATOR_POS_HORIZ_REL: "ESTIMATOR_POS_HORIZ_REL",
	ESTIMATOR_POS_HORIZ_ABS: "ESTIMATOR_POS_HORIZ_ABS",
	ESTIMATOR_POS_VERT_ABS: "ESTIMATOR_POS_VERT_ABS",
	ESTIMATOR_POS_VERT_AGL: "ESTIMATOR_POS_VERT_AGL",
	ESTIMATOR_CONST_POS_MODE: "ESTIMATOR_CONST_POS_MODE",
	ESTIMATOR_PRED_POS_HORIZ_REL: "ESTIMATOR_PRED_POS_HORIZ_REL",
	ESTIMATOR_PRED_POS_HORIZ_ABS: "ESTIMATOR_PRED_POS_HORIZ_ABS",
	ESTIMATOR_GPS_GLITCH: "ESTIMATOR_GPS_GLITCH",
	ESTIMATOR_ACCEL_ERROR: "ESTIMATOR_ACCEL_ERROR",
}

var values_ESTIMATOR_STATUS_FLAGS = map[string]ESTIMATOR_STATUS_FLAGS{
	"ESTIMATOR_ATTITUDE": ESTIMATOR_ATTITUDE,
	"ESTIMATOR_VELOCITY_HORIZ": ESTIMATOR_VELOCITY_HORIZ,
	"ESTIMATOR_VELOCITY_VERT": ESTIMATOR_VELOCITY_VERT,
	"ESTIMATOR_POS_HORIZ_REL": ESTIMATOR_POS_HORIZ_REL,
	"ESTIMATOR_POS_HORIZ_ABS": ESTIMATOR_POS_HORIZ_ABS,
	"ESTIMATOR_POS_VERT_ABS": ESTIMATOR_POS_VERT_ABS,
	"ESTIMATOR_POS_VERT_AGL": ESTIMATOR_POS_VERT_AGL,
	"ESTIMATOR_CONST_POS_MODE": ESTIMATOR_CONST_POS_MODE,
	"ESTIMATOR_PRED_POS_HORIZ_REL": ESTIMATOR_PRED_POS_HORIZ_REL,
	"ESTIMATOR_PRED_POS_HORIZ_ABS": ESTIMATOR_PRED_POS_HORIZ_ABS,
	"ESTIMATOR_GPS_GLITCH": ESTIMATOR_GPS_GLITCH,
	"ESTIMATOR_ACCEL_ERROR": ESTIMATOR_ACCEL_ERROR,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ESTIMATOR_STATUS_FLAGS) MarshalText() ([]byte, error) {
	if e == 0 {
		return []byte("0"), nil
	}
	var names []string
	for i := 0; i < 12; i++ {
		mask := ESTIMATOR_STATUS_FLAGS(1 << i)
		if e&mask == mask {
			names = append(names, labels_ESTIMATOR_STATUS_FLAGS[mask])
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ESTIMATOR_STATUS_FLAGS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask ESTIMATOR_STATUS_FLAGS
	for _, label := range labels {
		if value, ok := values_ESTIMATOR_STATUS_FLAGS[label]; ok {
			mask |= value
		} else if value, err := strconv.Atoi(label); err == nil {
			mask |= ESTIMATOR_STATUS_FLAGS(value)
		} else {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e ESTIMATOR_STATUS_FLAGS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
