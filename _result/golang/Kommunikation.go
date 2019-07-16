package golang

type Kommunikation struct {
	KommunikationNr      int     `json:KommunikationNr`
	Adresse              Adresse `json:Adresse`
	Bezeichnung          string  `json:Bezeichnung`
	Kommunikationsangabe string  `json:Kommunikationsangabe`
}
