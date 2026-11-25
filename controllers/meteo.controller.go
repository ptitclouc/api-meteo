package controllers

import (
	"errors"
	"log"
	"net/http"

	"pokedex/helpers"
	"pokedex/models"
	"pokedex/services"
)

// ListPokemon affiche la liste des Pokémon de la première génération.
func ListPokemon(w http.ResponseWriter, r *http.Request) {
	pokemons, err := services.GetFirstGeneration()
	if err != nil {
		renderErrorSevice(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "list", pokemons)
}

// ShowPokemon affiche la fiche du Pokémon dont l'identifiant est dans l'URL.
func ShowPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := services.GetPokemon(r.PathValue("id"))
	if err != nil {
		renderErrorSevice(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "pokemon", pokemon)
}

// NotFound affiche la page 404 pour toute URL inconnue du site.
func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, http.StatusNotFound, "Cette page n'existe pas.")
}

// renderAPIError traduit une erreur connue en page d'erreur, sur la route même.
func renderErrorSevice(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrNotFound):
		helpers.RenderError(w, http.StatusNotFound, "Ce Pokémon n'existe pas.")
	case errors.Is(err, models.ErrBadRequest):
		helpers.RenderError(w, http.StatusBadRequest, "Numéro de Pokémon invalide.")
	case errors.Is(err, models.ErrTimeout), errors.Is(err, models.ErrUnavailable):
		log.Println("API Tyradex :", err)
		helpers.RenderError(w, http.StatusServiceUnavailable, "Service momentanément indisponible.")
	default:
		log.Println("Erreur inattendue :", err)
		helpers.RenderError(w, http.StatusInternalServerError, "Une erreur est survenue.")
	}
}
