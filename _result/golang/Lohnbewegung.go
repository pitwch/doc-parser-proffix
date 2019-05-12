package golang

type Lohnbewegung struct {
	LohnbewegungNr float64     `json:LohnbewegungNr`
	Mitarbeiter    Mitarbeiter `json:Mitarbeiter`
	Lohnart        Lohnart     `json:Lohnart`
	VonPeriode     float64     `json:VonPeriode`
	BisPeriode     float64     `json:BisPeriode`
	Verrechnet     bool        `json:Verrechnet`
	Ansatz         float64     `json:Ansatz`
	Auftrag        Auftrag     `json:Auftrag`
	Betrag         float64     `json:Betrag`
	Bezeichnung    string      `json:Bezeichnung`
	Menge          float64     `json:Menge`
	Steuercode     Steuercode  `json:Steuercode`
}
