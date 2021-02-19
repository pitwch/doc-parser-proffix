package golang

type Positionsart struct {
	PositionsartNr  int          `json:PositionsartNr`
	Bezeichnung     string       `json:Bezeichnung`
	Vorgabezeit     float64      `json:Vorgabezeit`
	LohnperiodePlus bool         `json:LohnperiodePlus`
	Waehrung        Waehrung     `json:Waehrung`
	Steuercode      Steuercode   `json:Steuercode`
	Ertragskonto    Konto        `json:Ertragskonto`
	Kostenstelle    Kostenstelle `json:Kostenstelle`
	Kostenart       Kostenart    `json:Kostenart`
	KeinRabatt      bool         `json:KeinRabatt`
	Ferien          bool         `json:Ferien`
	StdPreis        float64      `json:StdPreis`
	GueltigVon      string       `json:GueltigVon`
	GueltigBis      string       `json:GueltigBis`
}
