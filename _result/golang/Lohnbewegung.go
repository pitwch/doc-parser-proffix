package golang

type Lohnbewegung struct {
	LohnbewegungNr int         `json:LohnbewegungNr`
	Mitarbeiter    Mitarbeiter `json:Mitarbeiter`
	Lohnart        Lohnart     `json:Lohnart`
	VonPeriode     int         `json:VonPeriode`
	BisPeriode     int         `json:BisPeriode`
	Verrechnet     bool        `json:Verrechnet`
	Ansatz         float64     `json:Ansatz`
	Auftrag        Auftrag     `json:Auftrag`
	Betrag         float64     `json:Betrag`
	Bezeichnung    string      `json:Bezeichnung`
	Menge          float64     `json:Menge`
	Steuercode     Steuercode  `json:Steuercode`
}
