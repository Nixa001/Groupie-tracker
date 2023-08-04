package models

type Artist struct {
	ID              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    int      `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationUrl     string   `json:"locations"`
	ConcertDatesUrl string   `json:"concertDates"`
	RelationsUrl    string   `json:"relations"`
}

// type artist interface{
// 	id              int
// 	image           string
// 	name            string
// 	members         string
// 	creationDate    int
// 	firstAlbum      string
// 	locationUrl     string
// 	concertDates string
// 	relations    string
// }

type Location struct {
	ID       int      `json:"id"`
	Location []string `json:"locations"`
	DatesUrl string   `json:"dates"`
}

type Date struct {
	ID   int      `json:"id"`
	Date []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
type Index struct {
	Index []Location `json:"index"`
}

// type Locationss struct {
// 	Index struct {
// 		ID       int      `json:"id"`
// 		Location []string `json:"locations"`
// 	} `json:"index"`
// }
