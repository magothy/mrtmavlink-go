//autogenerated:yes
//nolint:revive,misspell,govet,lll
package magothy
// Hardware status sent by an onboard computer.
type MessageOnboardComputerStatus struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Time since system boot.
	Uptime uint32
	// Type of the onboard computer: 0: Mission computer primary, 1: Mission computer backup 1, 2: Mission computer backup 2, 3: Compute node, 4-5: Compute spares, 6-9: Payload computers.
	Type uint8
	// CPU usage on the component in percent (100 - idle). A value of UINT8_MAX implies the field is unused.
	CpuCores [8]uint8
	// Combined CPU usage as the last 10 slices of 100 MS (a histogram). This allows to identify spikes in load that max out the system, but only for a short amount of time. A value of UINT8_MAX implies the field is unused.
	CpuCombined [10]uint8
	// GPU usage on the component in percent (100 - idle). A value of UINT8_MAX implies the field is unused.
	GpuCores [4]uint8
	// Combined GPU usage as the last 10 slices of 100 MS (a histogram). This allows to identify spikes in load that max out the system, but only for a short amount of time. A value of UINT8_MAX implies the field is unused.
	GpuCombined [10]uint8
	// Temperature of the board. A value of INT8_MAX implies the field is unused.
	TemperatureBoard int8
	// Temperature of the CPU core. A value of INT8_MAX implies the field is unused.
	TemperatureCore [8]int8
	// Fan speeds. A value of INT16_MAX implies the field is unused.
	FanSpeed [4]int16
	// Amount of used RAM on the component system. A value of UINT32_MAX implies the field is unused.
	RamUsage uint32
	// Total amount of RAM on the component system. A value of UINT32_MAX implies the field is unused.
	RamTotal uint32
	// Storage type: 0: HDD, 1: SSD, 2: EMMC, 3: SD card (non-removable), 4: SD card (removable). A value of UINT32_MAX implies the field is unused.
	StorageType [4]uint32
	// Amount of used storage space on the component system. A value of UINT32_MAX implies the field is unused.
	StorageUsage [4]uint32
	// Total amount of storage space on the component system. A value of UINT32_MAX implies the field is unused.
	StorageTotal [4]uint32
	// Link type: 0-9: UART, 10-19: Wired network, 20-29: Wifi, 30-39: Point-to-point proprietary, 40-49: Mesh proprietary
	LinkType [6]uint32
	// Network traffic from the component system. A value of UINT32_MAX implies the field is unused.
	LinkTxRate [6]uint32
	// Network traffic to the component system. A value of UINT32_MAX implies the field is unused.
	LinkRxRate [6]uint32
	// Network capacity from the component system. A value of UINT32_MAX implies the field is unused.
	LinkTxMax [6]uint32
	// Network capacity to the component system. A value of UINT32_MAX implies the field is unused.
	LinkRxMax [6]uint32
}

// GetID implements the message.Message interface.
func (*MessageOnboardComputerStatus) GetID() uint32 {
	return 390
}
