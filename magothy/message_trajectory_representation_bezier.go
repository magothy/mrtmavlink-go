//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Describe a trajectory using an array of up-to 5 bezier control points in the local frame (MAV_FRAME_LOCAL_NED).
type MessageTrajectoryRepresentationBezier struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Number of valid control points (up-to 5 points are possible)
	ValidPoints uint8
	// X-coordinate of bezier control points. Set to NaN if not being used
	PosX [5]float32
	// Y-coordinate of bezier control points. Set to NaN if not being used
	PosY [5]float32
	// Z-coordinate of bezier control points. Set to NaN if not being used
	PosZ [5]float32
	// Bezier time horizon. Set to NaN if velocity/acceleration should not be incorporated
	Delta [5]float32
	// Yaw. Set to NaN for unchanged
	PosYaw [5]float32
}

// GetID implements the message.Message interface.
func (*MessageTrajectoryRepresentationBezier) GetID() uint32 {
	return 333
}
