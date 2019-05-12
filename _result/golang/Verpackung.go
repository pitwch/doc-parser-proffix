package golang

type Verpackung struct {
	VerpackungNr string  `json:VerpackungNr`
	Bezeichnung  string  `json:Bezeichnung`
	Beschreibung string  `json:Beschreibung`
	Laenge       float64 `json:Laenge`
	Breite       float64 `json:Breite`
	Hoehe        float64 `json:Hoehe`
	Wandstaerke  float64 `json:Wandstaerke`
	Gewicht      float64 `json:Gewicht`
	Datei        string  `json:Datei`
}
