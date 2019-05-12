package golang

type Felddefinition struct {
	Tabellenname   string  `json:Tabellenname`
	Feldname       string  `json:Feldname`
	Tabellentitel  string  `json:Tabellentitel`
	Feldtitel      string  `json:Feldtitel`
	Dezimalstellen float64 `json:Dezimalstellen`
	Datentyp       string  `json:Datentyp`
	Standardwert   string  `json:Standardwert`
	Laenge         float64 `json:Laenge`
	Auswahl        string  `json:Auswahl`
	Zusatztabelle  bool    `json:Zusatztabelle`
	Zusatzfeld     bool    `json:Zusatzfeld`
}
