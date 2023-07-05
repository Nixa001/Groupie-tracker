package handlers

import (
	"encoding/json"
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	FetchApi(w, r)
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	utils.GetJsons(w)
	if r.URL.Path != "/" {
		Handle404Error(w, r)
		return
	}
	tmpl.Execute(w, nil)
}
func HandleArtists(w http.ResponseWriter, r *http.Request) {
	FetchApi(w, r)
	tmpl := template.Must(template.ParseFiles("templates/artists.html"))
	utils.GetJsons(w)
	if r.URL.Path != "/artists" {
		Handle404Error(w, r)
		return
	}
	utils.DataExec = map[string]interface{}{"Artists": utils.Artists}
	tmpl.Execute(w, utils.DataExec)
}

func ViewArtist(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/viewArtist.html"))
	tabUrl := strings.Split(r.URL.String(), "/")
	id, err := strconv.Atoi(tabUrl[len(tabUrl)-1])

	if id < 1 || id > 52 || err != nil {
		Handle400Error(w, r)
		return
	}
	if id > 0 {
		id = id - 1
	}
	FetchApi(w, r)
	utils.GetJsons(w)
	var Dates models.Date
	var Relations models.Relation
	var Locations models.Location

	responseData := utils.GetJson(w, utils.Artists[id].ConcertDatesUrl)
	json.Unmarshal(responseData, &Dates)
	responseData = utils.GetJson(w, utils.Artists[id].RelationsUrl)
	json.Unmarshal(responseData, &Relations)
	responseData = utils.GetJson(w, utils.Artists[id].LocationUrl)
	json.Unmarshal(responseData, &Locations)

	utils.Artists[id].FirstAlbum = utils.FormatDate(utils.Artists[id].FirstAlbum)
	for i, str := range Dates.Date {
		Dates.Date[i] = utils.FormatDate(str)
	}
	for i, str := range Locations.Location {
		Locations.Location[i] = utils.FormatStr(str)
	}
	for location, val := range Relations.DatesLocations {
		for _, dates := range val {
			Relations.DatesLocations[location] = []string{utils.FormatDate(dates)}
		}
		location = utils.FormatStr(location)
	}
	tmpl.Execute(w, map[interface{}]interface{}{
		"Artists":   utils.Artists[id],
		"Dates":     Dates.Date,
		"Relations": Relations.DatesLocations,
		"Locations": Locations})
}

func FetchApi(w http.ResponseWriter, r *http.Request) {
	data, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		Handle500Error(w, r)
		return
	}
	responsedata, _ := ioutil.ReadAll(data.Body)
	json.Unmarshal(responsedata, &utils.Urls)
}

func Handle404Error(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusNotFound,
		"TextErr": "Page Not Found"}
	ErrorPages(w, utils.DataExec)
}

func Handle400Error(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusBadRequest,
		"TextErr": "Bad Request"}
	ErrorPages(w, utils.DataExec)
}

func Handle500Error(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusInternalServerError,
		"TextErr": "Internal Server Error"}
	ErrorPages(w, utils.DataExec)
}

func ErrorPages(w http.ResponseWriter, data map[string]interface{}) {
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	tmpl.Execute(w, data)
}

func ContainAlpha(str string) bool {
	for _, runeValue := range str {
		if runeValue >= 'a' && runeValue <= 'z' {
			return true
		}
	}
	return false
}
