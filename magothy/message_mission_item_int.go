//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Message encoding a mission item. This message is emitted to announce
// the presence of a mission item and to set a mission item on the system. The mission item can be either in x, y, z meters (type: LOCAL) or x:lat, y:lon, z:altitude. Local frame is Z-down, right handed (NED), global frame is Z-up, right handed (ENU). NaN or INT32_MAX may be used in float/integer params (respectively) to indicate optional/default values (e.g. to use the component's current latitude, yaw rather than a specific value). See also https://mavlink.io/en/services/mission.html.
type MessageMissionItemInt struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Waypoint ID (sequence number). Starts at zero. Increases monotonically for each waypoint, no gaps in the sequence (0,1,2,3,4).
	Seq uint16
	// The coordinate system of the waypoint.
	Frame MAV_FRAME `mavenum:"uint8"`
	// The scheduled action for the waypoint.
	Command MAV_CMD `mavenum:"uint16"`
	// false:0, true:1
	Current uint8
	// Autocontinue to next waypoint
	Autocontinue uint8
	// PARAM1, see MAV_CMD enum
	Param1 float32
	// PARAM2, see MAV_CMD enum
	Param2 float32
	// PARAM3, see MAV_CMD enum
	Param3 float32
	// PARAM4, see MAV_CMD enum
	Param4 float32
	// PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
	X int32
	// PARAM6 / y position: local: x position in meters * 1e4, global: longitude in degrees *10^7
	Y int32
	// PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame.
	Z float32
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageMissionItemInt) GetID() uint32 {
	return 73
}
