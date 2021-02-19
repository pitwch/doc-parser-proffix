package golang

type Feiertag struct {
	FeiertagNr int     `json:FeiertagNr`
	Datum      string  `json:Datum`
	Ort        string  `json:Ort`
	Stunden    float64 `json:Stunden`
}
