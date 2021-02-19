package golang

type Artikeltext struct {
	ArtikeltextNr int     `json:ArtikeltextNr`
	Artikel       Artikel `json:Artikel`
	Bezeichnung1  string  `json:Bezeichnung1`
	Bezeichnung2  string  `json:Bezeichnung2`
	Bezeichnung3  string  `json:Bezeichnung3`
	Bezeichnung4  string  `json:Bezeichnung4`
	Bezeichnung5  string  `json:Bezeichnung5`
	Sprache       Sprache `json:Sprache`
}
