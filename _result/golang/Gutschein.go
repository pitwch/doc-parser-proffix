package golang

type Gutschein struct {
	GutscheinNr      string      `json:GutscheinNr`
	Bezeichnung      string      `json:Bezeichnung`
	Anzahl           int         `json:Anzahl`
	AnzahlEingeloest int         `json:AnzahlEingeloest`
	AnzahlProAdresse int         `json:AnzahlProAdresse`
	AnzahlVerkauft   int         `json:AnzahlVerkauft`
	Betrag           float64     `json:Betrag`
	Buchungsart      Buchungsart `json:Buchungsart`
	GueltigBis       string      `json:GueltigBis`
	GueltigVon       string      `json:GueltigVon`
	Konto            Konto       `json:Konto`
	Mindestbetrag    float64     `json:Mindestbetrag`
	NichtInWebshop   bool        `json:NichtInWebshop`
	NichtKumulierbar bool        `json:NichtKumulierbar`
	NurWebshop       bool        `json:NurWebshop`
	Prozent          float64     `json:Prozent`
	Waehrung         Waehrung    `json:Waehrung`
}
