package routes

import (
	"meteo/controllers"
	"net/http"
)

func MeteoRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", controllers.ListVille)
	mux.HandleFunc("GET /ville", controllers.ShowCommune)
	mux.HandleFunc("GET /ville/{id}", controllers.ShowMeteo)
}
