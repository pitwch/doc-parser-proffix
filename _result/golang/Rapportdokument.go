package golang

type Rapportdokument struct {
	RapportdokumentNr int     `json:"RapportdokumentNr"`
	Rapport           Rapport `json:"Rapport"`
	Bemerkungen       string  `json:"Bemerkungen"`
	Bezeichnung       string  `json:"Bezeichnung"`
	DateiNr           string  `json:"DateiNr"`
	Datum             string  `json:"Datum"`
	Dokumentgruppe    string  `json:"Dokumentgruppe"`
	Modul             string  `json:"Modul"`
}
