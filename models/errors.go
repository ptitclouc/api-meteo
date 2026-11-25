package models

import "errors"

// Erreurs connues, partagées par les couches api, services et controllers.
var (
	ErrNotFound    = errors.New("ressource introuvable")
	ErrBadRequest  = errors.New("demande invalide")
	ErrTimeout     = errors.New("l'API n'a pas répondu à temps")
	ErrUnavailable = errors.New("API indisponible")
)
