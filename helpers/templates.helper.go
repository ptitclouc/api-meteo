package helpers

import (
	"bytes"
	"html/template"
	"log"
	"meteo/api"
	"net/http"
)

var templates *template.Template

// LoadTemplates charge tous les templates une seule fois, au démarrage.
// Les fonctions (icone) doivent être déclarées AVANT le parsing des fichiers.
func LoadTemplates() {
	funcs := template.FuncMap{"icone": api.Icone}
	templates = template.Must(template.New("").Funcs(funcs).ParseGlob("templates/*.html"))
}

// Render exécute le template dans un buffer, puis envoie la page et son code.
func Render(w http.ResponseWriter, status int, name string, data any) {
	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, name, data); err != nil {
		log.Printf("Erreur template (%s) : %v\n", name, err)
		RenderError(w, http.StatusInternalServerError, "Erreur chargement de la page")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	buf.WriteTo(w)
}

// ErrorPage : données affichées par le template d'erreur.
type ErrorPage struct {
	Code    int
	Message string
}

// RenderError affiche la page d'erreur commune avec le bon code HTTP.
func RenderError(w http.ResponseWriter, status int, message string) {
	Render(w, status, "error", ErrorPage{Code: status, Message: message})
}