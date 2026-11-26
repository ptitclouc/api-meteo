package api

import (
	"fmt"
	"meteo/models"
)

func GetMeteo(id int) (models.Meteo, error) {
	var commune models.Meteo
	err := get(fmt.Sprintf("/commune/%d", id), &commune)
	return commune, err
}

func label(models.Meteo) (models.Commune, error){
	if models.label == 0{
		
	}
}