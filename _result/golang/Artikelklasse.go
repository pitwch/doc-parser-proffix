package golang

type Artikelklasse struct {
	ArtikelklasseNr string `json:ArtikelklasseNr`
	Bezeichnung     string `json:Bezeichnung`
	Gs1Lebensmittel bool   `json:Gs1Lebensmittel`
	Geloescht       bool   `json:Geloescht`
}
