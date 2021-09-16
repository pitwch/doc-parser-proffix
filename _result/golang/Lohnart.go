package golang

type Lohnart struct {
	LohnartNr   float64 `json:"LohnartNr"`
	Bezeichnung string  `json:"Bezeichnung"`
	Geloescht   bool    `json:"Geloescht"`
}
