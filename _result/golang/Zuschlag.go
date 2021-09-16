package golang

type Zuschlag struct {
	ZuschlagNr   string       `json:"ZuschlagNr"`
	Bezeichnung  string       `json:"Bezeichnung"`
	BisBetrag    float64      `json:"BisBetrag"`
	Kostenart    Kostenart    `json:"Kostenart"`
	Konto        Konto        `json:"Konto"`
	Kostenstelle Kostenstelle `json:"Kostenstelle"`
	Steuercode   Steuercode   `json:"Steuercode"`
	Zuschlag     float64      `json:"Zuschlag"`
}
