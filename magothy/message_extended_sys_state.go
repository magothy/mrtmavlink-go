//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Provides state for additional features
type MessageExtendedSysState struct {
	// The VTOL state if applicable. Is set to MAV_VTOL_STATE_UNDEFINED if UAV is not in VTOL configuration.
	VtolState MAV_VTOL_STATE `mavenum:"uint8"`
	// The landed state. Is set to MAV_LANDED_STATE_UNDEFINED if landed state is unknown.
	LandedState MAV_LANDED_STATE `mavenum:"uint8"`
}

// GetID implements the message.Message interface.
func (*MessageExtendedSysState) GetID() uint32 {
	return 245
}
