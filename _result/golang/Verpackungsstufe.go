package golang

type Verpackungsstufe struct {
	VerpackungsstufeNr int    `json:VerpackungsstufeNr`
	Bezeichnung        string `json:Bezeichnung`
	Indikator          int    `json:Indikator`
}
