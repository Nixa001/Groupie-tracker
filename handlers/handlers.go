package handlers

import (
	"encoding/json"
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"io"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	FetchApi(w, r)
	utils.GetJsons(w)
	if r.URL.Path != "/" {
		Handle404Error(w)
		return
	}
	if r.Method != "GET" && r.Method != "POST" {
		Handle405Error(w)
		return
	}

	utils.DataExec = map[string]interface{}{"Artists": utils.Artists}

	tmpl.ExecuteTemplate(w, "header", utils.DataExec)
	tmpl.ExecuteTemplate(w, "index", utils.DataExec)

}

func HandleArtists(w http.ResponseWriter, r *http.Request) {
	FetchApi(w, r)
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	utils.GetJsons(w)
	if r.URL.Path != "/artists" {
		Handle404Error(w)
		return
	}
	if r.Method != "GET" && r.Method != "POST" {
		Handle405Error(w)
		return
	}
	utils.DataExec = map[string]interface{}{
		"Artists": utils.Artists,
	}
	if r.Method == "POST" {

		str := r.PostFormValue("browsers")
		DataExec := Recherche(w, str)
		tmpl.ExecuteTemplate(w, "header", utils.DataExec)
		tmpl.ExecuteTemplate(w, "artists", map[string]interface{}{"Artists": DataExec})
		return
	}
	tmpl.ExecuteTemplate(w, "artists", utils.DataExec)
	tmpl.ExecuteTemplate(w, "header", utils.DataExec)
}

func ViewArtist(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	tabUrl := strings.Split(r.URL.String(), "/")
	if len(tabUrl) > 3 {
		Handle404Error(w)
		return

	}

	id, err := strconv.Atoi(tabUrl[len(tabUrl)-1])

	if id < 1 || id > 52 || err != nil {
		Handle400Error(w)
		return
	}
	if id > 0 {
		id = id - 1
	}
	if r.Method != "GET" && r.Method != "POST" {
		Handle405Error(w)
		return
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
	}
	DataExec := map[interface{}]interface{}{
		"Artists":   utils.Artists[id],
		"Dates":     Dates.Date,
		"Relations": Relations.DatesLocations,
		"Locations": Locations}
	DataExec1 := map[interface{}]interface{}{
		"Artists":   utils.Artists,
		"Dates":     Dates.Date,
		"Relations": Relations.DatesLocations,
		"Locations": Locations}
	tmpl.ExecuteTemplate(w, "header", DataExec1)
	tmpl.ExecuteTemplate(w, "viewArtist", DataExec)

}

func Recherche(w http.ResponseWriter, str string) []models.Artist {
	result := []models.Artist{}
	for _, artist := range utils.Artists {
		if strings.HasPrefix(strings.ToLower(artist.Name), strings.ToLower(str)) || strings.HasPrefix(strings.ToLower(artist.FirstAlbum), strings.ToLower(str)) || strings.HasPrefix(strings.ToLower(strconv.Itoa(artist.CreationDate)), strings.ToLower(str)) || strings.HasPrefix(strings.ToLower(artist.FirstAlbum), strings.ToLower(str)) {
			result = append(result, artist)
		}
		for _, v := range artist.Members {
			if strings.Contains(strings.ToLower(v), strings.ToLower(str)) {
				result = append(result, artist)

			}
		}
	}
	return result
}

func FetchApi(w http.ResponseWriter, r *http.Request) {
	data, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		Handle500Error(w)
		return
	}
	responsedata, _ := io.ReadAll(data.Body)
	json.Unmarshal(responsedata, &utils.Urls)
}

func Handle404Error(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusNotFound,
		"TextErr": "Page Not Found"}
	ErrorPages(w, utils.DataExec)
}

func Handle405Error(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusMethodNotAllowed,
		"TextErr": "Method Not Allowed"}
	ErrorPages(w, utils.DataExec)
}

func Handle400Error(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusBadRequest,
		"TextErr": "Bad Request"}
	ErrorPages(w, utils.DataExec)
}

func Handle500Error(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusInternalServerError,
		"TextErr": "Internal Server Error"}
	ErrorPages(w, utils.DataExec)
}

func ErrorPages(w http.ResponseWriter, data map[string]interface{}) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	tmpl.ExecuteTemplate(w, "header", data)
	tmpl.ExecuteTemplate(w, "error", data)
}
