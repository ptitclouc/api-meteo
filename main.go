package main



func main() {
	helpers.LoadTemplates()
	router := routes.MainRouter()

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
