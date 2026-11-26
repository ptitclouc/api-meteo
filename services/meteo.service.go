package services

import (
	"meteo/api"
	"meteo/models"
)

func GetMeteo(code string) (models.Meteo, error) {
	c, err := api.GetCommune(code)
	if err != nil {
		return models.Meteo{}, err
	}
	if len(c.Centre.Position) < 2 {
		return models.Meteo{}, models.ErrNotFound
	}
	meteo, err := api.GetMeteo(c.Centre.Position[1], c.Centre.Position[0])
	if err != nil {
		return models.Meteo{}, err
	}
	
	return meteo, nil
}