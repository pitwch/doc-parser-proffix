package golang

type Teilleistungsgruppe struct {
	TeilleistungsgruppeNr int    `json:"TeilleistungsgruppeNr"`
	Bezeichnung           string `json:"Bezeichnung"`
	Gliederung            string `json:"Gliederung"`
}
