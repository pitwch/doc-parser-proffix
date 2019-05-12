package golang

type Teilleistungsgruppe struct {
	TeilleistungsgruppeNr float64 `json:TeilleistungsgruppeNr`
	Bezeichnung           string  `json:Bezeichnung`
	Gliederung            string  `json:Gliederung`
}
