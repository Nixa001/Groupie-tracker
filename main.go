package main

import (
	"fmt"
	handlers "groupie-tracker/handlers"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artists", handlers.HandleArtists)
	http.HandleFunc("/artists/", handlers.ViewArtist)
	fmt.Println("Server start on port :8080")
	http.ListenAndServe(":8080", nil)

}
