//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Request all parameters of this component. All parameters should be emitted in response (as PARAM_EXT_VALUE or PARAM_EXT_VALUE_TRIMMED messages - see field: trimmed).
type MessageParamExtRequestList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Request _TRIMMED variants of PARAM_EXT_ messages. Set to 1 if _TRIMMED message variants are supported, and 0 otherwise. This signals the recipient that _TRIMMED messages are supported by the sender (and should be used if supported by the recipient).
	Trimmed uint8 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageParamExtRequestList) GetID() uint32 {
	return 321
}
