package golang

type Land struct {
	LandNr      string   `json:"LandNr"`
	Bezeichnung string   `json:"Bezeichnung"`
	Vorwahl     string   `json:"Vorwahl"`
	Waehrung    Waehrung `json:"Waehrung"`
	SADCode     int      `json:"SADCode"`
	CodePost    string   `json:"CodePost"`
}
