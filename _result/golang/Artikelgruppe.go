package golang

type Artikelgruppe struct {
	ArtikelgruppeNr string `json:"ArtikelgruppeNr"`
	Bezeichnung     string `json:"Bezeichnung"`
	Geloescht       bool   `json:"Geloescht"`
}
