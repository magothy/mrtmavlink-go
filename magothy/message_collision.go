//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Information about a potential collision
type MessageCollision struct {
	// Collision data source
	Src MAV_COLLISION_SRC `mavenum:"uint8"`
	// Unique identifier, domain based on src field
	Id uint32
	// Action that is being taken to avoid this collision
	Action MAV_COLLISION_ACTION `mavenum:"uint8"`
	// How concerned the aircraft is about this collision
	ThreatLevel MAV_COLLISION_THREAT_LEVEL `mavenum:"uint8"`
	// Estimated time until collision occurs
	TimeToMinimumDelta float32
	// Closest vertical distance between vehicle and object
	AltitudeMinimumDelta float32
	// Closest horizontal distance between vehicle and object
	HorizontalMinimumDelta float32
}

// GetID implements the message.Message interface.
func (*MessageCollision) GetID() uint32 {
	return 247
}
