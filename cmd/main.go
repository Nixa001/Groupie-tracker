package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var artists []Artist
var locations Location
var dates Date
var relations Relation
var Urls = map[string]string{}
var DataExec = map[string]interface{}{}

func main() {
	fs := http.FileServer(http.Dir("../template/static"))
	http.Handle("/template/static/", http.StripPrefix("/template/static/", fs))
	fetchApi()

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/artists/", viewArtist)
	fmt.Println("Server start on port :8080")
	http.ListenAndServe(":8080", nil)

}

func fetchApi() {
	data, _ := http.Get("https://groupietrackers.herokuapp.com/api")
	responsedata, _ := ioutil.ReadAll(data.Body)
	json.Unmarshal(responsedata, &Urls)
}

func getJson(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	return responseData
}
