package golang

type Steuercode struct {
	SteuercodeNr          int        `json:"SteuercodeNr"`
	Bezeichnung           string     `json:"Bezeichnung"`
	Prozent               float64    `json:"Prozent"`
	InklusivMWST          bool       `json:"InklusivMWST"`
	Konto                 Konto      `json:"Konto"`
	GueltigVon            string     `json:"GueltigVon"`
	GueltigBis            string     `json:"GueltigBis"`
	Ersatzsteuercode      Steuercode `json:"Ersatzsteuercode"`
	Ankauf                bool       `json:"Ankauf"`
	Bezugsteuerkonto      Konto      `json:"Bezugsteuerkonto"`
	Investition           bool       `json:"Investition"`
	DienstleistungAusland bool       `json:"DienstleistungAusland"`
	NichtSkontoberechtigt bool       `json:"NichtSkontoberechtigt"`
	Saldosteuersatz       float64    `json:"Saldosteuersatz"`
	Margenbesteuerung     bool       `json:"Margenbesteuerung"`
	Subventionen          bool       `json:"Subventionen"`
	Umsatzsteuercode      Steuercode `json:"Umsatzsteuercode"`
	Vorsteuer             bool       `json:"Vorsteuer"`
	ZahlProzent           float64    `json:"ZahlProzent"`
}
