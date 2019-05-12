package golang

type Gutschein struct {
	GutscheinNr      string      `json:GutscheinNr`
	Bezeichnung      string      `json:Bezeichnung`
	Anzahl           float64     `json:Anzahl`
	AnzahlEingeloest float64     `json:AnzahlEingeloest`
	AnzahlProAdresse float64     `json:AnzahlProAdresse`
	AnzahlVerkauft   float64     `json:AnzahlVerkauft`
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
