//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Winch status.
type MessageWinchStatus struct {
	// Timestamp (synced to UNIX time or since system boot).
	TimeUsec uint64
	// Length of line released. NaN if unknown
	LineLength float32
	// Speed line is being released or retracted. Positive values if being released, negative values if being retracted, NaN if unknown
	Speed float32
	// Tension on the line. NaN if unknown
	Tension float32
	// Voltage of the battery supplying the winch. NaN if unknown
	Voltage float32
	// Current draw from the winch. NaN if unknown
	Current float32
	// Temperature of the motor. INT16_MAX if unknown
	Temperature int16
	// Status flags
	Status MAV_WINCH_STATUS_FLAG `mavenum:"uint32"`
}

// GetID implements the message.Message interface.
func (*MessageWinchStatus) GetID() uint32 {
	return 9005
}
