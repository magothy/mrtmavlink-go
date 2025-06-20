//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Information about a component. For camera components instead use CAMERA_INFORMATION, and for autopilots use AUTOPILOT_VERSION. Components including GCSes should consider supporting requests of this message via MAV_CMD_REQUEST_MESSAGE.
type MessageComponentInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// The type of metadata being requested.
	MetadataType COMP_METADATA_TYPE `mavenum:"uint32"`
	// Unique uid for this metadata which a gcs can use for created cached metadata and understanding whether it's cache it up to date or whether it needs to download new data.
	MetadataUid uint32
	// Component definition URI. If prefix mavlinkftp:// file is downloaded from vehicle over mavlink ftp protocol. If prefix http[s]:// file is downloaded over internet. Files are json format. They can end in .gz to indicate file is in gzip format.
	MetadataUri string `mavlen:"70"`
	// Unique uid for the translation file associated with the metadata.
	TranslationUid uint32
	// The translations for strings within the metadata file. If null the either do not exist or may be included in the metadata file itself. The translations specified here supercede any which may be in the metadata file itself. The uri format is the same as component_metadata_uri . Files are in Json Translation spec format. Empty string indicates no tranlsation file.
	TranslationUri string `mavlen:"70"`
}

// GetID implements the message.Message interface.
func (*MessageComponentInformation) GetID() uint32 {
	return 395
}
