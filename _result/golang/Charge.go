package golang

type Charge struct {
	ChargeNr        string `json:ChargeNr`
	Bezeichnung     string `json:Bezeichnung`
	Inaktiv         bool   `json:Inaktiv`
	Auslaufsdatum   string `json:Auslaufsdatum`
	Haltbarkeit     int    `json:Haltbarkeit`
	Herstelldatum   string `json:Herstelldatum`
	UrsprungsCharge Charge `json:UrsprungsCharge`
}
