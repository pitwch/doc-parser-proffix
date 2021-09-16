package golang

type Verpackung struct {
	VerpackungNr string  `json:"VerpackungNr"`
	Bezeichnung  string  `json:"Bezeichnung"`
	Beschreibung string  `json:"Beschreibung"`
	Laenge       int     `json:"Laenge"`
	Breite       int     `json:"Breite"`
	Hoehe        int     `json:"Hoehe"`
	Wandstaerke  int     `json:"Wandstaerke"`
	Gewicht      float64 `json:"Gewicht"`
	Datei        string  `json:"Datei"`
}
