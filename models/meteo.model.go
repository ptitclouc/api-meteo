package models

type Commune struct{
	Code string `json:"code"`
	Name string `json:"nom"`
	Centre Coordonnee `json: "centre"`
}

type Coordonnee struct{
	Position []float32 `json:"coordinates"`
}

