package routes

import (
	"net/http"

	"pokedex/controllers"
)

// pokemonRoutes enregistre les routes de la partie Pokémon.
func pokemonRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", controllers.recherche)
	mux.HandleFunc("GET /pokemon/{id}", controllers.ShowMeteoVille)
}
