package golang

type Artikeluntergruppe struct {
	ArtikeluntergruppeNr string `json:ArtikeluntergruppeNr`
	Bezeichnung          string `json:Bezeichnung`
	Geloescht            bool   `json:Geloescht`
}
