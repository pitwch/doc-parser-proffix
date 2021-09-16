package golang

type Dokumentstatus struct {
	DokumentstatusNr string `json:"DokumentstatusNr"`
	Bezeichnung      string `json:"Bezeichnung"`
	Sperren          bool   `json:"Sperren"`
}
