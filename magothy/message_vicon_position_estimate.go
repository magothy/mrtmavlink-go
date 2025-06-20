//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Global position estimate from a Vicon motion system source.
type MessageViconPositionEstimate struct {
	// Timestamp (UNIX time or time since system boot)
	Usec uint64
	// Global X position
	X float32
	// Global Y position
	Y float32
	// Global Z position
	Z float32
	// Roll angle
	Roll float32
	// Pitch angle
	Pitch float32
	// Yaw angle
	Yaw float32
	// Row-major representation of 6x6 pose cross-covariance matrix upper right triangle (states: x, y, z, roll, pitch, yaw; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [21]float32 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageViconPositionEstimate) GetID() uint32 {
	return 104
}
