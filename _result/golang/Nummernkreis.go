package golang

type Nummernkreis struct {
	NummernkreisNr string `json:"NummernkreisNr"`
	NrVon          int    `json:"NrVon"`
	NrBis          int    `json:"NrBis"`
	NrAktuell      int    `json:"NrAktuell"`
}
