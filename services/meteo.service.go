package services

import (
	"fmt"
	"meteo/api"
	"meteo/models"
)

func GetMeteo(Latitude float32, Longitude float32) (models.Meteo, error) {
	var MeteoCommune models.Meteo
	Path := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&daily=temperature_2m_max,temperature_2m_min,weather_code&timezone=auto&forecast", Latitude, Longitude)
	err := api.Get(Path, MeteoCommune)
	return MeteoCommune, err
}
