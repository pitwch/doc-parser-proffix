package golang

type Verpackungsstufe struct {
	VerpackungsstufeNr float64 `json:VerpackungsstufeNr`
	Bezeichnung        string  `json:Bezeichnung`
	Indikator          float64 `json:Indikator`
}
