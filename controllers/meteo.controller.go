package controllers

import (
	"errors"
	"log"
	"meteo/api"
	"meteo/helpers"
	"meteo/models"
	"meteo/services"
	"net/http"
)

func ListVille(w http.ResponseWriter, r *http.Request) {
	villes, err := api.GetListeCommune()
	if err != nil {
		renderErrorSevice(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "list", villes)
}

func ShowMeteo(w http.ResponseWriter, r *http.Request) {
	meteo, err := services.GetMeteo(r.PathValue("id"))
	if err != nil {
		renderErrorSevice(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "meteo", meteo)
}

// NotFound affiche la page 404 pour toute URL inconnue du site.
func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, http.StatusNotFound, "Cette page n'existe pas.")
}

// renderAPIError traduit une erreur connue en page d'erreur, sur la route même.
func renderErrorSevice(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrNotFound):
		helpers.RenderError(w, http.StatusNotFound, "Cette commune n'existe pas")
	case errors.Is(err, models.ErrBadRequest):
		helpers.RenderError(w, http.StatusBadRequest, "invalide.")
	case errors.Is(err, models.ErrTimeout), errors.Is(err, models.ErrUnavailable):
		log.Println("API:", err)
		helpers.RenderError(w, http.StatusServiceUnavailable, "Service momentanément indisponible.")
	default:
		log.Println("Erreur inattendue :", err)
		helpers.RenderError(w, http.StatusInternalServerError, "Une erreur est survenue.")
	}
}
