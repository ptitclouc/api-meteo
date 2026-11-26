package services

import (
	"meteo/api"
	"meteo/models"
	"strconv"
)



func GetCommune(idParam string) (models.Commune, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return models.Commune{}, models.ErrBadRequest
	}
	return api.GetCommune(id)
}

func GetMeteo(idParam string) (models.Meteo, error)