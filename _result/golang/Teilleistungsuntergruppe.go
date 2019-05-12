package golang

type Teilleistungsuntergruppe struct {
	TeilleistungsuntergruppeNr float64 `json:TeilleistungsuntergruppeNr`
	Bezeichnung                string  `json:Bezeichnung`
	Gliederung                 string  `json:Gliederung`
}
