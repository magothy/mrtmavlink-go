//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Set the system mode, as defined by enum MAV_MODE. There is no target component id as the mode is by definition for the overall aircraft, not only for one component.
type MessageSetMode struct {
	// The system setting the mode
	TargetSystem uint8
	// The new base mode.
	BaseMode MAV_MODE `mavenum:"uint8"`
	// The new autopilot-specific mode. This field can be ignored by an autopilot.
	CustomMode uint32
}

// GetID implements the message.Message interface.
func (*MessageSetMode) GetID() uint32 {
	return 11
}
