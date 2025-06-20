//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package magothy

import (
	"strconv"
	"fmt"
)
// Type of AIS vessel, enum duplicated from AIS standard, https://gpsd.gitlab.io/gpsd/AIVDM.html
type AIS_TYPE uint64

const (
	// Not available (default).
	AIS_TYPE_UNKNOWN AIS_TYPE = 0
	AIS_TYPE_RESERVED_1 AIS_TYPE = 1
	AIS_TYPE_RESERVED_2 AIS_TYPE = 2
	AIS_TYPE_RESERVED_3 AIS_TYPE = 3
	AIS_TYPE_RESERVED_4 AIS_TYPE = 4
	AIS_TYPE_RESERVED_5 AIS_TYPE = 5
	AIS_TYPE_RESERVED_6 AIS_TYPE = 6
	AIS_TYPE_RESERVED_7 AIS_TYPE = 7
	AIS_TYPE_RESERVED_8 AIS_TYPE = 8
	AIS_TYPE_RESERVED_9 AIS_TYPE = 9
	AIS_TYPE_RESERVED_10 AIS_TYPE = 10
	AIS_TYPE_RESERVED_11 AIS_TYPE = 11
	AIS_TYPE_RESERVED_12 AIS_TYPE = 12
	AIS_TYPE_RESERVED_13 AIS_TYPE = 13
	AIS_TYPE_RESERVED_14 AIS_TYPE = 14
	AIS_TYPE_RESERVED_15 AIS_TYPE = 15
	AIS_TYPE_RESERVED_16 AIS_TYPE = 16
	AIS_TYPE_RESERVED_17 AIS_TYPE = 17
	AIS_TYPE_RESERVED_18 AIS_TYPE = 18
	AIS_TYPE_RESERVED_19 AIS_TYPE = 19
	// Wing In Ground effect.
	AIS_TYPE_WIG AIS_TYPE = 20
	AIS_TYPE_WIG_HAZARDOUS_A AIS_TYPE = 21
	AIS_TYPE_WIG_HAZARDOUS_B AIS_TYPE = 22
	AIS_TYPE_WIG_HAZARDOUS_C AIS_TYPE = 23
	AIS_TYPE_WIG_HAZARDOUS_D AIS_TYPE = 24
	AIS_TYPE_WIG_RESERVED_1 AIS_TYPE = 25
	AIS_TYPE_WIG_RESERVED_2 AIS_TYPE = 26
	AIS_TYPE_WIG_RESERVED_3 AIS_TYPE = 27
	AIS_TYPE_WIG_RESERVED_4 AIS_TYPE = 28
	AIS_TYPE_WIG_RESERVED_5 AIS_TYPE = 29
	AIS_TYPE_FISHING AIS_TYPE = 30
	AIS_TYPE_TOWING AIS_TYPE = 31
	// Towing: length exceeds 200m or breadth exceeds 25m.
	AIS_TYPE_TOWING_LARGE AIS_TYPE = 32
	// Dredging or other underwater ops.
	AIS_TYPE_DREDGING AIS_TYPE = 33
	AIS_TYPE_DIVING AIS_TYPE = 34
	AIS_TYPE_MILITARY AIS_TYPE = 35
	AIS_TYPE_SAILING AIS_TYPE = 36
	AIS_TYPE_PLEASURE AIS_TYPE = 37
	AIS_TYPE_RESERVED_20 AIS_TYPE = 38
	AIS_TYPE_RESERVED_21 AIS_TYPE = 39
	// High Speed Craft.
	AIS_TYPE_HSC AIS_TYPE = 40
	AIS_TYPE_HSC_HAZARDOUS_A AIS_TYPE = 41
	AIS_TYPE_HSC_HAZARDOUS_B AIS_TYPE = 42
	AIS_TYPE_HSC_HAZARDOUS_C AIS_TYPE = 43
	AIS_TYPE_HSC_HAZARDOUS_D AIS_TYPE = 44
	AIS_TYPE_HSC_RESERVED_1 AIS_TYPE = 45
	AIS_TYPE_HSC_RESERVED_2 AIS_TYPE = 46
	AIS_TYPE_HSC_RESERVED_3 AIS_TYPE = 47
	AIS_TYPE_HSC_RESERVED_4 AIS_TYPE = 48
	AIS_TYPE_HSC_UNKNOWN AIS_TYPE = 49
	AIS_TYPE_PILOT AIS_TYPE = 50
	// Search And Rescue vessel.
	AIS_TYPE_SAR AIS_TYPE = 51
	AIS_TYPE_TUG AIS_TYPE = 52
	AIS_TYPE_PORT_TENDER AIS_TYPE = 53
	// Anti-pollution equipment.
	AIS_TYPE_ANTI_POLLUTION AIS_TYPE = 54
	AIS_TYPE_LAW_ENFORCEMENT AIS_TYPE = 55
	AIS_TYPE_SPARE_LOCAL_1 AIS_TYPE = 56
	AIS_TYPE_SPARE_LOCAL_2 AIS_TYPE = 57
	AIS_TYPE_MEDICAL_TRANSPORT AIS_TYPE = 58
	// Noncombatant ship according to RR Resolution No. 18.
	AIS_TYPE_NONECOMBATANT AIS_TYPE = 59
	AIS_TYPE_PASSENGER AIS_TYPE = 60
	AIS_TYPE_PASSENGER_HAZARDOUS_A AIS_TYPE = 61
	AIS_TYPE_PASSENGER_HAZARDOUS_B AIS_TYPE = 62
	AIS_TYPE_AIS_TYPE_PASSENGER_HAZARDOUS_C AIS_TYPE = 63
	AIS_TYPE_PASSENGER_HAZARDOUS_D AIS_TYPE = 64
	AIS_TYPE_PASSENGER_RESERVED_1 AIS_TYPE = 65
	AIS_TYPE_PASSENGER_RESERVED_2 AIS_TYPE = 66
	AIS_TYPE_PASSENGER_RESERVED_3 AIS_TYPE = 67
	AIS_TYPE_AIS_TYPE_PASSENGER_RESERVED_4 AIS_TYPE = 68
	AIS_TYPE_PASSENGER_UNKNOWN AIS_TYPE = 69
	AIS_TYPE_CARGO AIS_TYPE = 70
	AIS_TYPE_CARGO_HAZARDOUS_A AIS_TYPE = 71
	AIS_TYPE_CARGO_HAZARDOUS_B AIS_TYPE = 72
	AIS_TYPE_CARGO_HAZARDOUS_C AIS_TYPE = 73
	AIS_TYPE_CARGO_HAZARDOUS_D AIS_TYPE = 74
	AIS_TYPE_CARGO_RESERVED_1 AIS_TYPE = 75
	AIS_TYPE_CARGO_RESERVED_2 AIS_TYPE = 76
	AIS_TYPE_CARGO_RESERVED_3 AIS_TYPE = 77
	AIS_TYPE_CARGO_RESERVED_4 AIS_TYPE = 78
	AIS_TYPE_CARGO_UNKNOWN AIS_TYPE = 79
	AIS_TYPE_TANKER AIS_TYPE = 80
	AIS_TYPE_TANKER_HAZARDOUS_A AIS_TYPE = 81
	AIS_TYPE_TANKER_HAZARDOUS_B AIS_TYPE = 82
	AIS_TYPE_TANKER_HAZARDOUS_C AIS_TYPE = 83
	AIS_TYPE_TANKER_HAZARDOUS_D AIS_TYPE = 84
	AIS_TYPE_TANKER_RESERVED_1 AIS_TYPE = 85
	AIS_TYPE_TANKER_RESERVED_2 AIS_TYPE = 86
	AIS_TYPE_TANKER_RESERVED_3 AIS_TYPE = 87
	AIS_TYPE_TANKER_RESERVED_4 AIS_TYPE = 88
	AIS_TYPE_TANKER_UNKNOWN AIS_TYPE = 89
	AIS_TYPE_OTHER AIS_TYPE = 90
	AIS_TYPE_OTHER_HAZARDOUS_A AIS_TYPE = 91
	AIS_TYPE_OTHER_HAZARDOUS_B AIS_TYPE = 92
	AIS_TYPE_OTHER_HAZARDOUS_C AIS_TYPE = 93
	AIS_TYPE_OTHER_HAZARDOUS_D AIS_TYPE = 94
	AIS_TYPE_OTHER_RESERVED_1 AIS_TYPE = 95
	AIS_TYPE_OTHER_RESERVED_2 AIS_TYPE = 96
	AIS_TYPE_OTHER_RESERVED_3 AIS_TYPE = 97
	AIS_TYPE_OTHER_RESERVED_4 AIS_TYPE = 98
	AIS_TYPE_OTHER_UNKNOWN AIS_TYPE = 99
)

var labels_AIS_TYPE = map[AIS_TYPE]string{
	AIS_TYPE_UNKNOWN: "AIS_TYPE_UNKNOWN",
	AIS_TYPE_RESERVED_1: "AIS_TYPE_RESERVED_1",
	AIS_TYPE_RESERVED_2: "AIS_TYPE_RESERVED_2",
	AIS_TYPE_RESERVED_3: "AIS_TYPE_RESERVED_3",
	AIS_TYPE_RESERVED_4: "AIS_TYPE_RESERVED_4",
	AIS_TYPE_RESERVED_5: "AIS_TYPE_RESERVED_5",
	AIS_TYPE_RESERVED_6: "AIS_TYPE_RESERVED_6",
	AIS_TYPE_RESERVED_7: "AIS_TYPE_RESERVED_7",
	AIS_TYPE_RESERVED_8: "AIS_TYPE_RESERVED_8",
	AIS_TYPE_RESERVED_9: "AIS_TYPE_RESERVED_9",
	AIS_TYPE_RESERVED_10: "AIS_TYPE_RESERVED_10",
	AIS_TYPE_RESERVED_11: "AIS_TYPE_RESERVED_11",
	AIS_TYPE_RESERVED_12: "AIS_TYPE_RESERVED_12",
	AIS_TYPE_RESERVED_13: "AIS_TYPE_RESERVED_13",
	AIS_TYPE_RESERVED_14: "AIS_TYPE_RESERVED_14",
	AIS_TYPE_RESERVED_15: "AIS_TYPE_RESERVED_15",
	AIS_TYPE_RESERVED_16: "AIS_TYPE_RESERVED_16",
	AIS_TYPE_RESERVED_17: "AIS_TYPE_RESERVED_17",
	AIS_TYPE_RESERVED_18: "AIS_TYPE_RESERVED_18",
	AIS_TYPE_RESERVED_19: "AIS_TYPE_RESERVED_19",
	AIS_TYPE_WIG: "AIS_TYPE_WIG",
	AIS_TYPE_WIG_HAZARDOUS_A: "AIS_TYPE_WIG_HAZARDOUS_A",
	AIS_TYPE_WIG_HAZARDOUS_B: "AIS_TYPE_WIG_HAZARDOUS_B",
	AIS_TYPE_WIG_HAZARDOUS_C: "AIS_TYPE_WIG_HAZARDOUS_C",
	AIS_TYPE_WIG_HAZARDOUS_D: "AIS_TYPE_WIG_HAZARDOUS_D",
	AIS_TYPE_WIG_RESERVED_1: "AIS_TYPE_WIG_RESERVED_1",
	AIS_TYPE_WIG_RESERVED_2: "AIS_TYPE_WIG_RESERVED_2",
	AIS_TYPE_WIG_RESERVED_3: "AIS_TYPE_WIG_RESERVED_3",
	AIS_TYPE_WIG_RESERVED_4: "AIS_TYPE_WIG_RESERVED_4",
	AIS_TYPE_WIG_RESERVED_5: "AIS_TYPE_WIG_RESERVED_5",
	AIS_TYPE_FISHING: "AIS_TYPE_FISHING",
	AIS_TYPE_TOWING: "AIS_TYPE_TOWING",
	AIS_TYPE_TOWING_LARGE: "AIS_TYPE_TOWING_LARGE",
	AIS_TYPE_DREDGING: "AIS_TYPE_DREDGING",
	AIS_TYPE_DIVING: "AIS_TYPE_DIVING",
	AIS_TYPE_MILITARY: "AIS_TYPE_MILITARY",
	AIS_TYPE_SAILING: "AIS_TYPE_SAILING",
	AIS_TYPE_PLEASURE: "AIS_TYPE_PLEASURE",
	AIS_TYPE_RESERVED_20: "AIS_TYPE_RESERVED_20",
	AIS_TYPE_RESERVED_21: "AIS_TYPE_RESERVED_21",
	AIS_TYPE_HSC: "AIS_TYPE_HSC",
	AIS_TYPE_HSC_HAZARDOUS_A: "AIS_TYPE_HSC_HAZARDOUS_A",
	AIS_TYPE_HSC_HAZARDOUS_B: "AIS_TYPE_HSC_HAZARDOUS_B",
	AIS_TYPE_HSC_HAZARDOUS_C: "AIS_TYPE_HSC_HAZARDOUS_C",
	AIS_TYPE_HSC_HAZARDOUS_D: "AIS_TYPE_HSC_HAZARDOUS_D",
	AIS_TYPE_HSC_RESERVED_1: "AIS_TYPE_HSC_RESERVED_1",
	AIS_TYPE_HSC_RESERVED_2: "AIS_TYPE_HSC_RESERVED_2",
	AIS_TYPE_HSC_RESERVED_3: "AIS_TYPE_HSC_RESERVED_3",
	AIS_TYPE_HSC_RESERVED_4: "AIS_TYPE_HSC_RESERVED_4",
	AIS_TYPE_HSC_UNKNOWN: "AIS_TYPE_HSC_UNKNOWN",
	AIS_TYPE_PILOT: "AIS_TYPE_PILOT",
	AIS_TYPE_SAR: "AIS_TYPE_SAR",
	AIS_TYPE_TUG: "AIS_TYPE_TUG",
	AIS_TYPE_PORT_TENDER: "AIS_TYPE_PORT_TENDER",
	AIS_TYPE_ANTI_POLLUTION: "AIS_TYPE_ANTI_POLLUTION",
	AIS_TYPE_LAW_ENFORCEMENT: "AIS_TYPE_LAW_ENFORCEMENT",
	AIS_TYPE_SPARE_LOCAL_1: "AIS_TYPE_SPARE_LOCAL_1",
	AIS_TYPE_SPARE_LOCAL_2: "AIS_TYPE_SPARE_LOCAL_2",
	AIS_TYPE_MEDICAL_TRANSPORT: "AIS_TYPE_MEDICAL_TRANSPORT",
	AIS_TYPE_NONECOMBATANT: "AIS_TYPE_NONECOMBATANT",
	AIS_TYPE_PASSENGER: "AIS_TYPE_PASSENGER",
	AIS_TYPE_PASSENGER_HAZARDOUS_A: "AIS_TYPE_PASSENGER_HAZARDOUS_A",
	AIS_TYPE_PASSENGER_HAZARDOUS_B: "AIS_TYPE_PASSENGER_HAZARDOUS_B",
	AIS_TYPE_AIS_TYPE_PASSENGER_HAZARDOUS_C: "AIS_TYPE_AIS_TYPE_PASSENGER_HAZARDOUS_C",
	AIS_TYPE_PASSENGER_HAZARDOUS_D: "AIS_TYPE_PASSENGER_HAZARDOUS_D",
	AIS_TYPE_PASSENGER_RESERVED_1: "AIS_TYPE_PASSENGER_RESERVED_1",
	AIS_TYPE_PASSENGER_RESERVED_2: "AIS_TYPE_PASSENGER_RESERVED_2",
	AIS_TYPE_PASSENGER_RESERVED_3: "AIS_TYPE_PASSENGER_RESERVED_3",
	AIS_TYPE_AIS_TYPE_PASSENGER_RESERVED_4: "AIS_TYPE_AIS_TYPE_PASSENGER_RESERVED_4",
	AIS_TYPE_PASSENGER_UNKNOWN: "AIS_TYPE_PASSENGER_UNKNOWN",
	AIS_TYPE_CARGO: "AIS_TYPE_CARGO",
	AIS_TYPE_CARGO_HAZARDOUS_A: "AIS_TYPE_CARGO_HAZARDOUS_A",
	AIS_TYPE_CARGO_HAZARDOUS_B: "AIS_TYPE_CARGO_HAZARDOUS_B",
	AIS_TYPE_CARGO_HAZARDOUS_C: "AIS_TYPE_CARGO_HAZARDOUS_C",
	AIS_TYPE_CARGO_HAZARDOUS_D: "AIS_TYPE_CARGO_HAZARDOUS_D",
	AIS_TYPE_CARGO_RESERVED_1: "AIS_TYPE_CARGO_RESERVED_1",
	AIS_TYPE_CARGO_RESERVED_2: "AIS_TYPE_CARGO_RESERVED_2",
	AIS_TYPE_CARGO_RESERVED_3: "AIS_TYPE_CARGO_RESERVED_3",
	AIS_TYPE_CARGO_RESERVED_4: "AIS_TYPE_CARGO_RESERVED_4",
	AIS_TYPE_CARGO_UNKNOWN: "AIS_TYPE_CARGO_UNKNOWN",
	AIS_TYPE_TANKER: "AIS_TYPE_TANKER",
	AIS_TYPE_TANKER_HAZARDOUS_A: "AIS_TYPE_TANKER_HAZARDOUS_A",
	AIS_TYPE_TANKER_HAZARDOUS_B: "AIS_TYPE_TANKER_HAZARDOUS_B",
	AIS_TYPE_TANKER_HAZARDOUS_C: "AIS_TYPE_TANKER_HAZARDOUS_C",
	AIS_TYPE_TANKER_HAZARDOUS_D: "AIS_TYPE_TANKER_HAZARDOUS_D",
	AIS_TYPE_TANKER_RESERVED_1: "AIS_TYPE_TANKER_RESERVED_1",
	AIS_TYPE_TANKER_RESERVED_2: "AIS_TYPE_TANKER_RESERVED_2",
	AIS_TYPE_TANKER_RESERVED_3: "AIS_TYPE_TANKER_RESERVED_3",
	AIS_TYPE_TANKER_RESERVED_4: "AIS_TYPE_TANKER_RESERVED_4",
	AIS_TYPE_TANKER_UNKNOWN: "AIS_TYPE_TANKER_UNKNOWN",
	AIS_TYPE_OTHER: "AIS_TYPE_OTHER",
	AIS_TYPE_OTHER_HAZARDOUS_A: "AIS_TYPE_OTHER_HAZARDOUS_A",
	AIS_TYPE_OTHER_HAZARDOUS_B: "AIS_TYPE_OTHER_HAZARDOUS_B",
	AIS_TYPE_OTHER_HAZARDOUS_C: "AIS_TYPE_OTHER_HAZARDOUS_C",
	AIS_TYPE_OTHER_HAZARDOUS_D: "AIS_TYPE_OTHER_HAZARDOUS_D",
	AIS_TYPE_OTHER_RESERVED_1: "AIS_TYPE_OTHER_RESERVED_1",
	AIS_TYPE_OTHER_RESERVED_2: "AIS_TYPE_OTHER_RESERVED_2",
	AIS_TYPE_OTHER_RESERVED_3: "AIS_TYPE_OTHER_RESERVED_3",
	AIS_TYPE_OTHER_RESERVED_4: "AIS_TYPE_OTHER_RESERVED_4",
	AIS_TYPE_OTHER_UNKNOWN: "AIS_TYPE_OTHER_UNKNOWN",
}

var values_AIS_TYPE = map[string]AIS_TYPE{
	"AIS_TYPE_UNKNOWN": AIS_TYPE_UNKNOWN,
	"AIS_TYPE_RESERVED_1": AIS_TYPE_RESERVED_1,
	"AIS_TYPE_RESERVED_2": AIS_TYPE_RESERVED_2,
	"AIS_TYPE_RESERVED_3": AIS_TYPE_RESERVED_3,
	"AIS_TYPE_RESERVED_4": AIS_TYPE_RESERVED_4,
	"AIS_TYPE_RESERVED_5": AIS_TYPE_RESERVED_5,
	"AIS_TYPE_RESERVED_6": AIS_TYPE_RESERVED_6,
	"AIS_TYPE_RESERVED_7": AIS_TYPE_RESERVED_7,
	"AIS_TYPE_RESERVED_8": AIS_TYPE_RESERVED_8,
	"AIS_TYPE_RESERVED_9": AIS_TYPE_RESERVED_9,
	"AIS_TYPE_RESERVED_10": AIS_TYPE_RESERVED_10,
	"AIS_TYPE_RESERVED_11": AIS_TYPE_RESERVED_11,
	"AIS_TYPE_RESERVED_12": AIS_TYPE_RESERVED_12,
	"AIS_TYPE_RESERVED_13": AIS_TYPE_RESERVED_13,
	"AIS_TYPE_RESERVED_14": AIS_TYPE_RESERVED_14,
	"AIS_TYPE_RESERVED_15": AIS_TYPE_RESERVED_15,
	"AIS_TYPE_RESERVED_16": AIS_TYPE_RESERVED_16,
	"AIS_TYPE_RESERVED_17": AIS_TYPE_RESERVED_17,
	"AIS_TYPE_RESERVED_18": AIS_TYPE_RESERVED_18,
	"AIS_TYPE_RESERVED_19": AIS_TYPE_RESERVED_19,
	"AIS_TYPE_WIG": AIS_TYPE_WIG,
	"AIS_TYPE_WIG_HAZARDOUS_A": AIS_TYPE_WIG_HAZARDOUS_A,
	"AIS_TYPE_WIG_HAZARDOUS_B": AIS_TYPE_WIG_HAZARDOUS_B,
	"AIS_TYPE_WIG_HAZARDOUS_C": AIS_TYPE_WIG_HAZARDOUS_C,
	"AIS_TYPE_WIG_HAZARDOUS_D": AIS_TYPE_WIG_HAZARDOUS_D,
	"AIS_TYPE_WIG_RESERVED_1": AIS_TYPE_WIG_RESERVED_1,
	"AIS_TYPE_WIG_RESERVED_2": AIS_TYPE_WIG_RESERVED_2,
	"AIS_TYPE_WIG_RESERVED_3": AIS_TYPE_WIG_RESERVED_3,
	"AIS_TYPE_WIG_RESERVED_4": AIS_TYPE_WIG_RESERVED_4,
	"AIS_TYPE_WIG_RESERVED_5": AIS_TYPE_WIG_RESERVED_5,
	"AIS_TYPE_FISHING": AIS_TYPE_FISHING,
	"AIS_TYPE_TOWING": AIS_TYPE_TOWING,
	"AIS_TYPE_TOWING_LARGE": AIS_TYPE_TOWING_LARGE,
	"AIS_TYPE_DREDGING": AIS_TYPE_DREDGING,
	"AIS_TYPE_DIVING": AIS_TYPE_DIVING,
	"AIS_TYPE_MILITARY": AIS_TYPE_MILITARY,
	"AIS_TYPE_SAILING": AIS_TYPE_SAILING,
	"AIS_TYPE_PLEASURE": AIS_TYPE_PLEASURE,
	"AIS_TYPE_RESERVED_20": AIS_TYPE_RESERVED_20,
	"AIS_TYPE_RESERVED_21": AIS_TYPE_RESERVED_21,
	"AIS_TYPE_HSC": AIS_TYPE_HSC,
	"AIS_TYPE_HSC_HAZARDOUS_A": AIS_TYPE_HSC_HAZARDOUS_A,
	"AIS_TYPE_HSC_HAZARDOUS_B": AIS_TYPE_HSC_HAZARDOUS_B,
	"AIS_TYPE_HSC_HAZARDOUS_C": AIS_TYPE_HSC_HAZARDOUS_C,
	"AIS_TYPE_HSC_HAZARDOUS_D": AIS_TYPE_HSC_HAZARDOUS_D,
	"AIS_TYPE_HSC_RESERVED_1": AIS_TYPE_HSC_RESERVED_1,
	"AIS_TYPE_HSC_RESERVED_2": AIS_TYPE_HSC_RESERVED_2,
	"AIS_TYPE_HSC_RESERVED_3": AIS_TYPE_HSC_RESERVED_3,
	"AIS_TYPE_HSC_RESERVED_4": AIS_TYPE_HSC_RESERVED_4,
	"AIS_TYPE_HSC_UNKNOWN": AIS_TYPE_HSC_UNKNOWN,
	"AIS_TYPE_PILOT": AIS_TYPE_PILOT,
	"AIS_TYPE_SAR": AIS_TYPE_SAR,
	"AIS_TYPE_TUG": AIS_TYPE_TUG,
	"AIS_TYPE_PORT_TENDER": AIS_TYPE_PORT_TENDER,
	"AIS_TYPE_ANTI_POLLUTION": AIS_TYPE_ANTI_POLLUTION,
	"AIS_TYPE_LAW_ENFORCEMENT": AIS_TYPE_LAW_ENFORCEMENT,
	"AIS_TYPE_SPARE_LOCAL_1": AIS_TYPE_SPARE_LOCAL_1,
	"AIS_TYPE_SPARE_LOCAL_2": AIS_TYPE_SPARE_LOCAL_2,
	"AIS_TYPE_MEDICAL_TRANSPORT": AIS_TYPE_MEDICAL_TRANSPORT,
	"AIS_TYPE_NONECOMBATANT": AIS_TYPE_NONECOMBATANT,
	"AIS_TYPE_PASSENGER": AIS_TYPE_PASSENGER,
	"AIS_TYPE_PASSENGER_HAZARDOUS_A": AIS_TYPE_PASSENGER_HAZARDOUS_A,
	"AIS_TYPE_PASSENGER_HAZARDOUS_B": AIS_TYPE_PASSENGER_HAZARDOUS_B,
	"AIS_TYPE_AIS_TYPE_PASSENGER_HAZARDOUS_C": AIS_TYPE_AIS_TYPE_PASSENGER_HAZARDOUS_C,
	"AIS_TYPE_PASSENGER_HAZARDOUS_D": AIS_TYPE_PASSENGER_HAZARDOUS_D,
	"AIS_TYPE_PASSENGER_RESERVED_1": AIS_TYPE_PASSENGER_RESERVED_1,
	"AIS_TYPE_PASSENGER_RESERVED_2": AIS_TYPE_PASSENGER_RESERVED_2,
	"AIS_TYPE_PASSENGER_RESERVED_3": AIS_TYPE_PASSENGER_RESERVED_3,
	"AIS_TYPE_AIS_TYPE_PASSENGER_RESERVED_4": AIS_TYPE_AIS_TYPE_PASSENGER_RESERVED_4,
	"AIS_TYPE_PASSENGER_UNKNOWN": AIS_TYPE_PASSENGER_UNKNOWN,
	"AIS_TYPE_CARGO": AIS_TYPE_CARGO,
	"AIS_TYPE_CARGO_HAZARDOUS_A": AIS_TYPE_CARGO_HAZARDOUS_A,
	"AIS_TYPE_CARGO_HAZARDOUS_B": AIS_TYPE_CARGO_HAZARDOUS_B,
	"AIS_TYPE_CARGO_HAZARDOUS_C": AIS_TYPE_CARGO_HAZARDOUS_C,
	"AIS_TYPE_CARGO_HAZARDOUS_D": AIS_TYPE_CARGO_HAZARDOUS_D,
	"AIS_TYPE_CARGO_RESERVED_1": AIS_TYPE_CARGO_RESERVED_1,
	"AIS_TYPE_CARGO_RESERVED_2": AIS_TYPE_CARGO_RESERVED_2,
	"AIS_TYPE_CARGO_RESERVED_3": AIS_TYPE_CARGO_RESERVED_3,
	"AIS_TYPE_CARGO_RESERVED_4": AIS_TYPE_CARGO_RESERVED_4,
	"AIS_TYPE_CARGO_UNKNOWN": AIS_TYPE_CARGO_UNKNOWN,
	"AIS_TYPE_TANKER": AIS_TYPE_TANKER,
	"AIS_TYPE_TANKER_HAZARDOUS_A": AIS_TYPE_TANKER_HAZARDOUS_A,
	"AIS_TYPE_TANKER_HAZARDOUS_B": AIS_TYPE_TANKER_HAZARDOUS_B,
	"AIS_TYPE_TANKER_HAZARDOUS_C": AIS_TYPE_TANKER_HAZARDOUS_C,
	"AIS_TYPE_TANKER_HAZARDOUS_D": AIS_TYPE_TANKER_HAZARDOUS_D,
	"AIS_TYPE_TANKER_RESERVED_1": AIS_TYPE_TANKER_RESERVED_1,
	"AIS_TYPE_TANKER_RESERVED_2": AIS_TYPE_TANKER_RESERVED_2,
	"AIS_TYPE_TANKER_RESERVED_3": AIS_TYPE_TANKER_RESERVED_3,
	"AIS_TYPE_TANKER_RESERVED_4": AIS_TYPE_TANKER_RESERVED_4,
	"AIS_TYPE_TANKER_UNKNOWN": AIS_TYPE_TANKER_UNKNOWN,
	"AIS_TYPE_OTHER": AIS_TYPE_OTHER,
	"AIS_TYPE_OTHER_HAZARDOUS_A": AIS_TYPE_OTHER_HAZARDOUS_A,
	"AIS_TYPE_OTHER_HAZARDOUS_B": AIS_TYPE_OTHER_HAZARDOUS_B,
	"AIS_TYPE_OTHER_HAZARDOUS_C": AIS_TYPE_OTHER_HAZARDOUS_C,
	"AIS_TYPE_OTHER_HAZARDOUS_D": AIS_TYPE_OTHER_HAZARDOUS_D,
	"AIS_TYPE_OTHER_RESERVED_1": AIS_TYPE_OTHER_RESERVED_1,
	"AIS_TYPE_OTHER_RESERVED_2": AIS_TYPE_OTHER_RESERVED_2,
	"AIS_TYPE_OTHER_RESERVED_3": AIS_TYPE_OTHER_RESERVED_3,
	"AIS_TYPE_OTHER_RESERVED_4": AIS_TYPE_OTHER_RESERVED_4,
	"AIS_TYPE_OTHER_UNKNOWN": AIS_TYPE_OTHER_UNKNOWN,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e AIS_TYPE) MarshalText() ([]byte, error) {
	if name, ok := labels_AIS_TYPE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *AIS_TYPE) UnmarshalText(text []byte) error {
	if value, ok := values_AIS_TYPE[string(text)]; ok {
	   *e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
	   *e = AIS_TYPE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e AIS_TYPE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
