//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Emit an encrypted signature / key identifying this system. PLEASE NOTE: This protocol has been kept simple, so transmitting the key requires an encrypted channel for true safety.
type MessageAuthKey struct {
	// key
	Key string `mavlen:"32"`
}

// GetID implements the message.Message interface.
func (*MessageAuthKey) GetID() uint32 {
	return 7
}
