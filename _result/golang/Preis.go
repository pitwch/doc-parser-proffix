package golang

type Preis struct {
	Bezeichnung   string       `json:Bezeichnung`
	Preisstaffel  Preisstaffel `json:Preisstaffel`
	Preiscode     int          `json:Preiscode`
	StaffelVon    float64      `json:StaffelVon`
	StaffelBis    float64      `json:StaffelBis`
	Wert          float64      `json:Wert`
	InProzent     bool         `json:InProzent`
	IstRabatt     bool         `json:IstRabatt`
	Preis         float64      `json:Preis`
	Waehrung      Waehrung     `json:Waehrung`
	Zusatzartikel Artikel      `json:Zusatzartikel`
}
