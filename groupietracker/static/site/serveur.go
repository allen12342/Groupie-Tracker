package groupietracker

// Artist est la structure pour stocker les informations sur les artistes
type Artist struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Group    string `json:"group"`
	Location string `json:"location"`
	Albums   []Album
	Dates    []string
}

// Album est la structure pour stocker les informations sur les albums
type Album struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

// Location est la structure pour stocker les informations sur les lieux
type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getLocations() []Location {
	return []Location{
		{ID: 1, Name: "Lieu 1"},
		{ID: 2, Name: "Lieu 2"},
	}
}

func getArtists() []Artist {
	return []Artist{
		{
			ID:       1,
			Name:     "Artiste 1",
			Group:    "Groupe 1",
			Location: "Lieu 1",
			Albums: []Album{
				{ID: 1, Name: "Album 1", Artist: "Artiste 1"},
				{ID: 2, Name: "Album 2", Artist: "Artiste 1"},
			},
			Dates: []string{"2022-05-01", "2022-06-02"},
		},
		{
			ID:       2,
			Name:     "Artiste 2",
			Group:    "Groupe 2",
			Location: "Lieu 2",
			Albums: []Album{
				{ID: 1, Name: "Album 1", Artist: "Artiste 2"},
				{ID: 2, Name: "Album 2", Artist: "Artiste 2"},
			},
			Dates: []string{"2022-07-01", "2022-08-02"},
		},
	}
}
