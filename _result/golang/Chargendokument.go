package golang

type Chargendokument struct {
	DokumentNr     int    `json:"DokumentNr"`
	ChargeNr       string `json:"ChargeNr"`
	Bemerkung      string `json:"Bemerkung"`
	Bezeichnung    string `json:"Bezeichnung"`
	DateiNr        string `json:"DateiNr"`
	Datum          string `json:"Datum"`
	Dokumentgruppe string `json:"Dokumentgruppe"`
}
