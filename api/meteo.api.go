package api

func Icone(codeVMO int) string {
	switch {
	case codeVMO == 0:
		return "☀️"
	case codeVMO == 1:
		return "🌤️"
	case codeVMO == 2:
		return "⛅"
	case codeVMO == 3:
		return "☁️"
	case codeVMO == 45 || codeVMO == 48:
		return "🌫️"
	case codeVMO >= 51 && codeVMO <= 57:
		return "🌦️"
	case codeVMO >= 61 && codeVMO <= 67:
		return "🌧️"
	case codeVMO >= 71 && codeVMO <= 77:
		return "🌨️"
	case codeVMO >= 80 && codeVMO <= 82:
		return "🌦️"
	}
	return ""
}
