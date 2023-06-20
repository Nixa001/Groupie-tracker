package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Group struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func main() {
	group := Group{
		ID:           1,
		Image:        "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
		Name:         "Queen",
		Members:      []string{"Freddie Mercury", "Brian May", "John Daecon", "Roger Meddows-Taylor", "Mike Grose", "Barry Mitchell", "Doug Fogie"},
		CreationDate: 1970,
		FirstAlbum:   "14-12-1973",
		Locations:    "https://groupietrackers.herokuapp.com/api/locations/1",
		ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/1",
		Relations:    "https://groupietrackers.herokuapp.com/api/relation/1",
	}

	// Effectuer la requête GET pour récupérer le contenu de l'URL
	response, err := http.Get(group.Locations)
	if err != nil {
		fmt.Println("Erreur lors de la requête GET:", err)
		return
	}
	defer response.Body.Close()

	// Lire le contenu de la réponse
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		return
	}

	// Afficher le contenu de l'URL
	fmt.Println("Contenu de l'URL", group.Locations, ":")
	fmt.Println(string(body))
}
