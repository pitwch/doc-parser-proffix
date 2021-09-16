package golang

type Stundenart struct {
	StundenartNr               string  `json:"StundenartNr"`
	Bezeichnung                string  `json:"Bezeichnung"`
	ZeitErfassen               bool    `json:"ZeitErfassen"`
	Vorgabezeit                float64 `json:"Vorgabezeit"`
	Stundenantrag              bool    `json:"Stundenantrag"`
	Ferien                     bool    `json:"Ferien"`
	IstStundenPausen           bool    `json:"IstStundenPausen"`
	IstStundenKuerzungen       bool    `json:"IstStundenKuerzungen"`
	KeineStundenabrechnung     bool    `json:"KeineStundenabrechnung"`
	InOutlookKalenderSpeichern bool    `json:"InOutlookKalenderSpeichern"`
	LohnperiodePlus1           bool    `json:"LohnperiodePlus1"`
	GueltigVon                 string  `json:"GueltigVon"`
	GueltigBis                 string  `json:"GueltigBis"`
	Lohnart                    Lohnart `json:"Lohnart"`
	FarbeKalender              int     `json:"FarbeKalender"`
}
