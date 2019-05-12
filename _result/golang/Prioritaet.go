package golang

type Prioritaet struct {
	PrioritaetNr         int     `json:PrioritaetNr`
	Bezeichnung          string  `json:Bezeichnung`
	Feiertage            bool    `json:Feiertage`
	Loesungszeit         float64 `json:Loesungszeit`
	EskalationsmeldungAn string  `json:EskalationsmeldungAn`
	NurEineWarnung       bool    `json:NurEineWarnung`
	Reaktionszeit        float64 `json:Reaktionszeit`
	Reihenfolge          float64 `json:Reihenfolge`
	Samstag              bool    `json:Samstag`
	Sonntag              bool    `json:Sonntag`
	WarnungEskalation    float64 `json:WarnungEskalation`
	Zeitbereich1Bis      string  `json:Zeitbereich1Bis`
	Zeitbereich1Von      string  `json:Zeitbereich1Von`
	Zeitbereich2Bis      string  `json:Zeitbereich2Bis`
	Zeitbereich2Von      string  `json:Zeitbereich2Von`
	ZeitenRechnen        int     `json:ZeitenRechnen`
}
