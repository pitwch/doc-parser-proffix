package golang

type Felddefinition struct {
	Tabellenname       string `json:Tabellenname`
	Feldname           string `json:Feldname`
	Tabellentitel      string `json:Tabellentitel`
	Feldtitel          string `json:Feldtitel`
	Dezimalstellen     int    `json:Dezimalstellen`
	Datentyp           string `json:Datentyp`
	Standardwert       string `json:Standardwert`
	Laenge             int    `json:Laenge`
	Auswahl            string `json:Auswahl`
	Zusatztabelle      bool   `json:Zusatztabelle`
	Zusatzfeld         bool   `json:Zusatzfeld`
	NichtKopieren      bool   `json:NichtKopieren`
	ZusatzfeldDatentyp string `json:ZusatzfeldDatentyp`
	Pflicht            bool   `json:Pflicht`
}
