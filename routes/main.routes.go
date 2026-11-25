package routes

import (
	"net/http"

	"pokedex/controllers"
)

// MainRouter crée le routeur principal et y ajoute toutes les routes du site.
func MainRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Fichiers statiques : /static/css/style.css → assets/css/style.css
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	MeteoRoutes(mux)

	// Toute autre URL : page 404 du site
	mux.HandleFunc("/", controllers.NotFound)
	return mux
}