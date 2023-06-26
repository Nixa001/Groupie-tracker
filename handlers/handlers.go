package handlers

import (
	"encoding/json"
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
	utils.GetJsons()
	if r.URL.Path != "/" {
		Handle404Error(w, r)
		return
	}
	tmpl.Execute(w, nil)
}
func HandleArtists(w http.ResponseWriter, r *http.Request) {
	FetchApi(w, r)
	tmpl := template.Must(template.ParseFiles("templates/artists.html"))
	utils.GetJsons()
	if r.URL.Path != "/artists" {
		Handle404Error(w, r)
		return
	}
	utils.DataExec = map[string]interface{}{"Artists": utils.Artists}
	tmpl.Execute(w, utils.DataExec)
}

func ViewArtist(w http.ResponseWriter, r *http.Request) {
	tabUrl := strings.Split(r.URL.String(), "/")
	id, err := strconv.Atoi(tabUrl[len(tabUrl)-1])
	if err != nil {
		Handle500Error(w, r)
	}
	if id < 1 || id > 52 {
		Handle400Error(w, r)
		return
	}
	if id > 0 {
		id = id - 1
	}
	FetchApi(w, r)
	utils.GetJsons()
	responseData := utils.GetJson(utils.Artists[id].ConcertDatesUrl)
	json.Unmarshal(responseData, &utils.Dates)
	responseData = utils.GetJson(utils.Artists[id].RelationsUrl)
	json.Unmarshal(responseData, &utils.Relations)
	responseData = utils.GetJson(utils.Artists[id].LocationUrl)

	json.Unmarshal(responseData, &utils.Locations)
	tmpl := template.Must(template.ParseFiles("templates/viewArtist.html"))
	utils.Artists[id].FirstAlbum = utils.FormatDate(utils.Artists[id].FirstAlbum)
	for i, str := range utils.Dates.Date {
		utils.Dates.Date[i] = utils.FormatDate(str)
	}
	for i, str := range utils.Locations.Location {
		utils.Locations.Location[i] = utils.FormatStr(str)
	}
	// var count int
	// for location, val := range utils.Relations.DatesLocations {
	// 	for i := count; i < len(val); i++ {
	// 		utils.Relations.DatesLocations[location] = []string{utils.FormatDate(utils.Relations.DatesLocations[location][i])}
	// 		fmt.Println(utils.Relations.DatesLocations[location])
	// 		break
	// 	}
	// 	count++
	// }
	// fmt.Println(utils.Dates.Date)
	tmpl.Execute(w, map[interface{}]interface{}{
		"Artists":   utils.Artists[id],
		"Dates":     utils.Dates.Date,
		"Relations": utils.Relations.DatesLocations,
		"Locations": utils.Locations})
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
		"TextErr": "Page Not Fount"}
	ErrorPages(w, r, utils.DataExec)
}

func Handle400Error(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusBadRequest,
		"TextErr": "Bad Request"}
	ErrorPages(w, r, utils.DataExec)
}

func Handle500Error(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusInternalServerError,
		"TextErr": "Internal Server Error"}
	ErrorPages(w, r, utils.DataExec)
}

func ErrorPages(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	tmpl.Execute(w, data)
}
