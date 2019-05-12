package golang

type Lieferart struct {
	LieferartNr  string       `json:LieferartNr`
	Bezeichnung  string       `json:Bezeichnung`
	Betrag       float64      `json:Betrag`
	Steuercode   Steuercode   `json:Steuercode`
	Konto        Konto        `json:Konto`
	Kostenstelle Kostenstelle `json:Kostenstelle`
	Kostenart    Kostenart    `json:Kostenart`
	Inaktiv      bool         `json:Inaktiv`
}
