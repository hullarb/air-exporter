// based on https://github.com/marco-hoyer/zcan/blob/master/src/main/python/zcan/mapping.py
package main

import "math"

type comfoairFrame struct {
	name          string
	unit          string
	dataConverter func([8]byte) float64
}

func data2Temp(data [8]byte) float64 {
	if len(data) < 2 {
		return math.MaxFloat64
	}
	return (float64(data[0]) - float64(data[1])) / 10
}

func data2Vol(data [8]byte) float64 {
	if len(data) < 2 {
		return math.MaxFloat64
	}
	return float64(data[0]) + float64(data[1])*255
}

func byte2Float(data [8]byte) float64 {
	if len(data) < 1 {
		return math.MaxFloat64
	}
	return float64(data[0])
}

var frameIDMap = map[uint32]comfoairFrame{
	0x0010C041: {
		name:          "comfocool_profile",
		unit:          "0:Normal, 1:Cool, 2:Warm",
		dataConverter: byte2Float,
	},

	0x00148041: {
		name:          "unknown_decreasing_number",
		unit:          "unknown",
		dataConverter: byte2Float,
	},
	0x00454041: {
		name:          "temperature_inlet_before_recuperator",
		unit:          "°C",
		dataConverter: data2Temp,
	},
	0x00458041: {
		name:          "temperature_inlet_after_recuperator",
		unit:          "°C",
		dataConverter: data2Temp,
	},
	0x001E0041: {
		name:          "air_volume_input_ventilator",
		unit:          "m3",
		dataConverter: data2Vol,
	},
	0x001DC041: {
		name:          "air_volume_output_ventilator",
		unit:          "m3",
		dataConverter: data2Vol,
	},
	0x001E8041: {
		name:          "speed_input_ventilator",
		unit:          "rpm",
		dataConverter: data2Vol,
	},
	0x001E4041: {
		name:          "speed_output_ventilator",
		unit:          "rpm",
		dataConverter: data2Vol,
	},
	0x00488041: {
		name:          "air_humidity_outlet_before_recuperator",
		unit:          "%",
		dataConverter: byte2Float,
	},
	0x0048C041: {
		name:          "air_humidity_outlet_after_recuperator",
		unit:          "%",
		dataConverter: byte2Float,
	},
	0x00200041: {
		name:          "total_power_consumption",
		unit:          "W",
		dataConverter: byte2Float,
	},
	0x001D4041: {
		name:          "power_percent_output_ventilator",
		unit:          "%",
		dataConverter: byte2Float,
	},
	0x001D8041: {
		name:          "power_percent_input_ventilator",
		unit:          "%",
		dataConverter: byte2Float,
	},
	0x0082C042: {
		name:          "0x0082C042",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x004C4041: {
		name:          "0x004C4041",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x00384041: {
		name:          "0x00384041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},
	0x00144041: {
		name:          "remaining_s_in_currend_ventilation_mode",
		unit:          "s",
		dataConverter: data2Vol,
	},
	0x00824042: {
		name:          "0x00824042",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x00810042: {
		name:          "0x00810042",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x00208041: {
		name:          "total_power_consumption_2",
		unit:          "kWh",
		dataConverter: data2Vol,
	},
	0x00344041: {
		name:          "0x00344041",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x00370041: {
		name:          "0x00370041",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x00300041: {
		name:          "days_until_next_filter_change",
		unit:          "days",
		dataConverter: data2Vol,
	},
	0x00044041: {
		name:          "0x00044041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},
	0x00204041: {
		name:          "total_power_consumption_this_year",
		unit:          "kWh",
		dataConverter: data2Vol,
	},
	0x00084041: {
		name:          "0x00084041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},
	0x00804042: {
		name:          "0x00804042",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x00644041: {
		name:          "0x00644041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},
	0x00354041: {
		name:          "0x00354041",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x00390041: {
		name:          "frost_disbalance",
		unit:          "%",
		dataConverter: byte2Float,
	},

	0x0035C041: {
		name:          "total_power_savings",
		unit:          "kWh",
		dataConverter: data2Vol,
	},

	0x0044C041: {
		name:          "temperature_outlet_after_recuperator",
		unit:          "°C",
		dataConverter: data2Temp,
	},

	0x0080C042: {
		name:          "0x0080C042",
		unit:          "unknown",
		dataConverter: data2Vol,
	},

	0x000E0041: {
		name:          "0x000E0041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00604041: {
		name:          "0x00604041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00450041: {
		name:          "0x00450041",
		unit:          "unknown",
		dataConverter: data2Vol,
	},

	0x00378041: {
		name:          "0x00378041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00818042: {
		name:          "0x00818042",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00820042: {
		name:          "0x00820042",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x001D0041: {
		name:          "0x001D0041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00490041: {
		name:          "0x00490041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00350041: {
		name:          "0x00350041",
		unit:          "unknown",
		dataConverter: data2Vol,
	},

	0x0081C042: {
		name:          "0x0081C042",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00448041: {
		name:          "temperature_outlet_before_recuperator",
		unit:          "°C",
		dataConverter: data2Temp,
	},

	0x00560041: {
		name:          "0x00560041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00374041: {
		name:          "0x00374041",
		unit:          "unknown",
		dataConverter: data2Vol,
	},

	0x00808042: {
		name:          "0x00808042",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00040041: {
		name:          "0x00040041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x10040001: {
		name:          "10040001",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00120041: {
		name:          "0x00120041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00688041: {
		name:          "0x00688041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00358041: {
		name:          "total_power_savings_this_year",
		unit:          "kWh",
		dataConverter: data2Vol,
	},

	0x00104041: {
		name:          "ventilation_level",
		unit:          "level",
		dataConverter: byte2Float,
	},

	0x00544041: {
		name:          "0x00544041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00498041: {
		name:          "air_humidity_inlet_after_recuperator",
		unit:          "%",
		dataConverter: byte2Float,
	},

	0x00814042: {
		name:          "0x00814042",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x000C4041: {
		name:          "0x000C4041",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00828042: {
		name:          "0x00828042",
		unit:          "unknown",
		dataConverter: byte2Float,
	},

	0x00494041: {
		name:          "air_humidity_inlet_before_recuperator",
		unit:          "%",
		dataConverter: byte2Float,
	},

	0x004C8041: {
		name:          "0x004C8041",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x00388041: {
		name:          "0x00388041",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x00188041: {
		name:          "bypass_a_status",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x00184041: {
		name:          "bypass_b_status",
		unit:          "unknown",
		dataConverter: data2Vol,
	},
	0x00108041: {
		name:          "bypass_state",
		unit:          "0=auto,1=open,2=close",
		dataConverter: byte2Float,
	},
	0x0038C041: {
		name:          "bypass_open",
		unit:          "%",
		dataConverter: byte2Float,
	},
}

// command_mapping = {
//     "set_ventilation_level_0": b'T1F07505180100201C00000000\r',
//     "set_ventilation_level_1": b'T1F07505180100201C00000100\r',
//     "set_ventilation_level_2": b'T1F07505180100201C00000200\r',
//     "set_ventilation_level_3": b'T1F07505180100201C00000300\r',
//     "auto_mode": b'T1F075051485150801\r', # verified (also: T1F051051485150801\r)
//     "manual_mode": b'T1F07505180084150101000000\r', # verified (also: T1F051051485150801\r)
//     "temperature_profile_cool": b'T0010C041101\r',
//     "temperature_profile_normal": b'T0010C041100\r',
//     "temperature_profile_warm": b'T0010C041102\r',
//     "close_bypass": b'T00108041102\r',
//     "open_bypass": b'T00108041101\r',
//     "auto_bypass": b'T00108041100\r',
//     "basis_menu": b"T00400041100\r",
//     "extended_menu": b"T00400041101\r"
// }
