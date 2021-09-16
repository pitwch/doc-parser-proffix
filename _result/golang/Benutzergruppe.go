package golang

type Benutzergruppe struct {
	BenutzergruppeNr string `json:"BenutzergruppeNr"`
	Bezeichnung      string `json:"Bezeichnung"`
	IstAdmin         bool   `json:"IstAdmin"`
	ADGruppe         string `json:"ADGruppe"`
}
