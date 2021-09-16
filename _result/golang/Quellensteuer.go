package golang

type Quellensteuer struct {
	QuellensteuerNr    int     `json:"QuellensteuerNr"`
	Bezeichnung        string  `json:"Bezeichnung"`
	JahrGueltigVon     int     `json:"JahrGueltigVon"`
	JahrGueltigBis     int     `json:"JahrGueltigBis"`
	Importdateiname    string  `json:"Importdateiname"`
	KeineUebermittlung bool    `json:"KeineUebermittlung"`
	Positionen         string  `json:"Positionen"`
	AbBetrag           float64 `json:"AbBetrag"`
	Prozent            float64 `json:"Prozent"`
}
