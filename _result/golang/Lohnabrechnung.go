package golang

type Lohnabrechnung struct {
	LohnabrechnungNr       int         `json:LohnabrechnungNr`
	Periode                int         `json:Periode`
	SubNr                  int         `json:SubNr`
	Buchungsdatum          string      `json:Buchungsdatum`
	Titel                  string      `json:Titel`
	Positionen             string      `json:Positionen`
	LohnbewegungPositionNr float64     `json:LohnbewegungPositionNr`
	Mitarbeiter            Mitarbeiter `json:Mitarbeiter`
	Lohnart                Lohnart     `json:Lohnart`
	Menge                  float64     `json:Menge`
	Ansatz                 float64     `json:Ansatz`
	Betrag                 float64     `json:Betrag`
}
