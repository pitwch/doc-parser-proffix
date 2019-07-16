package golang

type Artikeloption struct {
	ArtikeloptionNr int     `json:ArtikeloptionNr`
	Anzahl          float64 `json:Anzahl`
	Artikel         Artikel `json:Artikel`
	Bemerkungen     string  `json:Bemerkungen`
	BemerkungenRTF  string  `json:BemerkungenRTF`
	Option          Artikel `json:Option`
	Sprache         Sprache `json:Sprache`
}
