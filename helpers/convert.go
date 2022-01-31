package helpers

func BoolToInt(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}

func DeviceStatus(value string) int {
	switch value {
	case "INACTIVE":
		return 0
	case "ACTIVE":
		return 1
	case "CANCEL":
		return 2
	case "OUTDATED":
		return 3
	default:
		return 0
	}
}

func DeviceSimActive(value string) int {
	switch value {
	case "NOT_ASSIGN":
		return 0
	case "ASSIGN":
		return 1
	case "ASSIGNED":
		return 2
	default:
		return 0
	}
}

/* */
func AlertSettingTypeToInt(t string) int {
	switch t {
	case "MOVE":
		return 1
	case "NIGHT":
		return 2
	case "OFF_MOVE":
		return 3
	case "DISTANCE":
		return 4
	case "BATTERY_VOLTAGE":
		return 5
	case "NEAR_OWN_PHONE":
		return 6
	case "DEVICE_TURN_ON":
		return 7
	case "DEVICE_TURN_OFF":
		return 8
	case "DEVICE_CHANGE_PASS":
		return 9
	case "VEHICLE_ON":
		return 10
	case "VEHICLE_OFF":
		return 11
	case "DEVICE_OFF":
		return 12
	case "MISS_KEY":
		return 13
	default:
		return 0
	}
}

func AlertSettingIntToType(t int) string {
	switch t {
	case 1:
		return "MOVE"
	case 2:
		return "NIGHT"
	case 3:
		return "OFF_MOVE"
	case 4:
		return "DISTANCE"
	case 5:
		return "BATTERY_VOLTAGE"
	case 6:
		return "NEAR_OWN_PHONE"
	case 7:
		return "DEVICE_TURN_ON"
	case 8:
		return "DEVICE_TURN_OFF"
	case 9:
		return "DEVICE_CHANGE_PASS"
	case 10:
		return "VEHICLE_ON"
	case 11:
		return "VEHICLE_OFF"
	case 12:
		return "DEVICE_OFF"
	case 13:
		return "MISS_KEY"
	default:
		return "MOVE"
	}
}
