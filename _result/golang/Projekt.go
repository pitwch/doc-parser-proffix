package golang

type Projekt struct {
	ProjektNr      string  `json:ProjektNr`
	Bezeichnung    string  `json:Bezeichnung`
	Adresse        Adresse `json:Adresse`
	Teilleistungen string  `json:Teilleistungen`
}
