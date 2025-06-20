//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Set a parameter value. In order to deal with message loss (and retransmission of PARAM_EXT_SET), when setting a parameter value and the new value is the same as the current value, you will immediately get a PARAM_ACK_ACCEPTED response. If the current state is PARAM_ACK_IN_PROGRESS, you will accordingly receive a PARAM_ACK_IN_PROGRESS in response.
type MessageParamExtSet struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter value
	ParamValue string `mavlen:"128"`
	// Parameter type.
	ParamType MAV_PARAM_EXT_TYPE `mavenum:"uint8"`
}

// GetID implements the message.Message interface.
func (*MessageParamExtSet) GetID() uint32 {
	return 323
}
