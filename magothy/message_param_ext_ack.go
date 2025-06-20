//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Response from a PARAM_EXT_SET message.
type MessageParamExtAck struct {
	// Parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter value (new value if PARAM_ACK_ACCEPTED, current value otherwise)
	ParamValue string `mavlen:"128"`
	// Parameter type.
	ParamType MAV_PARAM_EXT_TYPE `mavenum:"uint8"`
	// Result code.
	ParamResult PARAM_ACK `mavenum:"uint8"`
}

// GetID implements the message.Message interface.
func (*MessageParamExtAck) GetID() uint32 {
	return 324
}
