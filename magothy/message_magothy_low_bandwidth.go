//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// The heartbeat message shows that a system or component is present and responding. The type and autopilot fields (along with the message component id), allow the receiving system to treat further messages from this system appropriately (e.g. by laying out the user interface based on the autopilot). This microservice is documented at https://mavlink.io/en/services/heartbeat.html
type MessageMagothyLowBandwidth struct {
	// Vehicle or component type. For a flight controller component the vehicle type (quadrotor, helicopter, etc.). For other components the component type (e.g. camera, gimbal, etc.). This should be used in preference to component id for identifying the component type.
	Type MAV_TYPE `mavenum:"uint8"`
	// A bitfield for use for autopilot-specific flags
	CustomMode uint32
	// Bitmap showing which onboard controllers and sensors are present. Value of 0: not present. Value of 1: present.
	OnboardControlSensorsPresent MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	// Bitmap showing which onboard controllers and sensors are enabled:  Value of 0: not enabled. Value of 1: enabled.
	OnboardControlSensorsEnabled MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	// Bitmap showing which onboard controllers and sensors have an error (or are operational). Value of 0: error. Value of 1: healthy.
	OnboardControlSensorsHealth MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	// Battery voltage, UINT16_MAX: Voltage not sent by autopilot
	VoltageBattery uint16
	// Battery current, -1: Current not sent by autopilot
	CurrentBattery int16
	// Battery energy remaining, -1: Battery remaining energy not sent by autopilot
	BatteryRemaining int8
	// Sequence number of the current active mission item. UINT16_MAX: not in mission
	MissionSeq uint16
	// Latitude (WGS84, EGM96 ellipsoid)
	Lat int32
	// Longitude (WGS84, EGM96 ellipsoid)
	Lon int32
	// GPS ground speed. If unknown, set to: UINT16_MAX
	Speed uint16
	// Course over ground in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	Course uint16
	// Number of satellites visible. If unknown, set to 255
	SatellitesVisible uint8
	// Heading in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	Heading uint16
	// 1 if position measurement is independent (gps), else 0
	IsPositionIndependent uint8
	// Position estimate error. If unknown, set to: UINT16_MAX
	PositionError uint16
	// Desired ground speed. If unknown, set to: UINT16_MAX
	DesiredSpeed uint16
	// Desired course over ground in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	DesiredCourse uint16
	// UUID of most recent mode change
	GcsSetModeUuidLsb uint32 `mavext:"true"`
	// CRC-16/CCITT-FALSE of serialized loaded mission
	MissionCrc uint16 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageMagothyLowBandwidth) GetID() uint32 {
	return 50004
}
