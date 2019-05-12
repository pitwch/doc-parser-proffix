package golang

type SerieNummer struct {
	SerieNummerNr float64    `json:SerieNummerNr`
	SerieNr       string     `json:SerieNr`
	Artikel       Artikel    `json:Artikel`
	Lagerort      Lagerort   `json:Lagerort`
	Lagerplatz    Lagerplatz `json:Lagerplatz`
	Charge        Charge     `json:Charge`
	Geloescht     bool       `json:Geloescht`
}
