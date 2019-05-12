package golang

type Kondition struct {
	KonditionNr    float64 `json:KonditionNr`
	Bezeichnung    string  `json:Bezeichnung`
	Inaktiv        bool    `json:Inaktiv`
	EndeJahr       bool    `json:EndeJahr`
	EndeMonat      bool    `json:EndeMonat`
	EndeQuartal    bool    `json:EndeQuartal`
	NettoTage      float64 `json:NettoTage`
	Skonto1Tage    float64 `json:Skonto1Tage`
	Skonto1Prozent float64 `json:Skonto1Prozent`
	Skonto2Tage    float64 `json:Skonto2Tage`
	Skonto2Prozent float64 `json:Skonto2Prozent`
}
