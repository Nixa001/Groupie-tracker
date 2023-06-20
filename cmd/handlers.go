package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../template/index.html"))
	getJsons()
	// if r.URL.Path != "/" {
	// 	DataExec = map[string]interface{}{"ErrNum": http.StatusNotFound, "TextErr": "Page not found"}
	// 	w.WriteHeader(http.StatusNotFound)
	// 	errorPages(w, r, DataExec)
	// 	return
	// }
	DataExec = map[string]interface{}{"Artists": artists}
	tmpl.Execute(w, DataExec)
}

func viewArtist(w http.ResponseWriter, r *http.Request) {
	tabUrl := strings.Split(r.URL.String(), "/")
	id, _ := strconv.Atoi(tabUrl[len(tabUrl)-1])
	if id > 0 {
		id = id - 1
	}
	// Images := artists[id].Image
	response, _ := http.Get((artists[id].Location))
	responseData, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseData, &locations)

	tmpl := template.Must(template.ParseFiles("../template/viewArtist.html"))
	tmpl.Execute(w, map[interface{}]interface{}{"Artists": artists[id], "Locations": locations})
}

func getJsons() {
	json.Unmarshal(getJson(Urls["artists"]), &artists)
	json.Unmarshal(getJson(Urls["locations"]), &locations)
	json.Unmarshal(getJson(Urls["relation"]), &relations)
	json.Unmarshal(getJson(Urls["dates"]), &dates)
}

func errorPages(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	tmpl := template.Must(template.ParseFiles("../template/error.html"))
	tmpl.Execute(w, data)
}
