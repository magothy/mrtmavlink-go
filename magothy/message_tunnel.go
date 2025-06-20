//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Message for transporting "arbitrary" variable-length data from one component to another (broadcast is not forbidden, but discouraged). The encoding of the data is usually extension specific, i.e. determined by the source, and is usually not documented as part of the MAVLink specification.
type MessageTunnel struct {
	// System ID (can be 0 for broadcast, but this is discouraged)
	TargetSystem uint8
	// Component ID (can be 0 for broadcast, but this is discouraged)
	TargetComponent uint8
	// A code that identifies the content of the payload (0 for unknown, which is the default). If this code is less than 32768, it is a 'registered' payload type and the corresponding code should be added to the MAV_TUNNEL_PAYLOAD_TYPE enum. Software creators can register blocks of types as needed. Codes greater than 32767 are considered local experiments and should not be checked in to any widely distributed codebase.
	PayloadType MAV_TUNNEL_PAYLOAD_TYPE `mavenum:"uint16"`
	// Length of the data transported in payload
	PayloadLength uint8
	// Variable length payload. The payload length is defined by payload_length. The entire content of this block is opaque unless you understand the encoding specified by payload_type.
	Payload [128]uint8
}

// GetID implements the message.Message interface.
func (*MessageTunnel) GetID() uint32 {
	return 385
}
