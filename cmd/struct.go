package main

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstDate"`
	Location     string   `json:"locations"`
}

type Locationss struct {
	Index struct {
		ID       int      `json:"id"`
		Location []string `json:"locations"`
	} `json:"index"`
}

type Location struct {
	ID       int      `json:"id"`
	Location []string `json:"locations"`
}

type Date struct {
	Index []struct {
		ID   int      `json:"id"`
		Date []string `json:"dates"`
	}
}

type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}
}
