package golang

type Charge struct {
	ChargeNr    string `json:ChargeNr`
	Bezeichnung string `json:Bezeichnung`
	Inaktiv     bool   `json:Inaktiv`
}
