package utils

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

var Artists []models.Artist
var Locations models.Location
var Dates models.Date
var Relations models.Relation
var DataExec = map[string]interface{}{}
var Urls = map[string]string{}

func GetJson(url string) []byte {
	response, _ := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	responseData, _ := ioutil.ReadAll(response.Body)
	return responseData
}
func GetJsons() {
	json.Unmarshal(GetJson(Urls["artists"]), &Artists)
	json.Unmarshal(GetJson(Urls["locations"]), &Locations)
	json.Unmarshal(GetJson(Urls["relation"]), &Relations)
	json.Unmarshal(GetJson(Urls["dates"]), &Dates)
}

func ErrorPages(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	tmpl.Execute(w, data)
}

func FormatDate(dateStr string) string {
	dateStr = strings.ReplaceAll(dateStr, "*", "")
	date, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		os.Exit(1)
	}
	formatedDate := date.Format("Jan 2, 2006")
	return formatedDate
}
func FormatStr(str string) string {
	str = strings.ReplaceAll(str, "-", " ")
	str = strings.ReplaceAll(str, "_", " ")
	str = strings.Title(str)
	return str
}
