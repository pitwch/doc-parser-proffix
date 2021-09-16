package golang

type Kontoklasse struct {
	KontoklasseNr string `json:"KontoklasseNr"`
	Bezeichnung   string `json:"Bezeichnung"`
	Kontotyp      int    `json:"Kontotyp"`
}
