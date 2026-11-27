package api

import (
	"fmt"
	"meteo/models"
)

const (
	meteoURL = "https://api.open-meteo.com"
)

func GetMeteo(Latitude float32, Longitude float32) (models.Meteo, error) {
	var meteo models.Meteo
	Path := fmt.Sprintf("%s/v1/forecast?latitude=%f&longitude=%f&daily=temperature_2m_max,temperature_2m_min,weather_code&timezone=auto", meteoURL, Latitude, Longitude)
	err := Get(Path, &meteo)
	return meteo, err
}

func Icone(codeVMO int) string {
	switch {
	case codeVMO == 0:
		return "Ciel dégagé ☀️"
	case codeVMO == 1:
		return "Peu nuageux 🌤️"
	case codeVMO == 2:
		return "Partiellement nuageux ⛅"
	case codeVMO == 3:
		return "Couvert☁️"
	case codeVMO == 45 || codeVMO == 48:
		return "Brouillard🌫️"
	case codeVMO >= 51 && codeVMO <= 57:
		return "Brouillard 🌦️"
	case codeVMO >= 61 && codeVMO <= 67:
		return "Bruine 🌧️"
	case codeVMO >= 71 && codeVMO <= 77:
		return "Pluie 🌨️"
	case codeVMO >= 80 && codeVMO <= 82:
		return "Averses 🌦️"
	case codeVMO == 85 || codeVMO == 86:
		return "Averse de neige🌨️"
	case codeVMO >= 95 && codeVMO <= 99:
		return "Orage ⛈️"
	}
	return ""
}
