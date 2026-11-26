package models

type Commune struct {
	Code   string     `json:"code"`
	Name   string     `json:"nom"`
	Centre Coordonnee `json: "centre"`
}

type Coordonnee struct {
	Position []float32 `json:"coordinates"`
}

type Meteo struct {
	Latitude   float32    `json: "latitude"`
	Longitude  float32    `json: "longitude"`
	Journalier semaine    `json: "daily"`
	Unités     UnitésUtil `json: "daily_units"`
}

type semaine struct {
	Date    []string `json: "time"`
	TempMax []float32 `json: "temperature_2m_max"`
	TempMin []float32`json: "temperature_2m_min"`
	Icone   []string `json: "weather_code"`
}

type UnitésUtil struct {
	TMaxDeg string `json: "temperature_2m_max"`
	TMinDeg string `json: "temperature_2m_min"`
}
