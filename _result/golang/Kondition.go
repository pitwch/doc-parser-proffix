package golang

type Kondition struct {
	KonditionNr    int     `json:"KonditionNr"`
	Bezeichnung    string  `json:"Bezeichnung"`
	Inaktiv        bool    `json:"Inaktiv"`
	EndeJahr       bool    `json:"EndeJahr"`
	EndeMonat      bool    `json:"EndeMonat"`
	EndeQuartal    bool    `json:"EndeQuartal"`
	NettoTage      int     `json:"NettoTage"`
	Skonto1Tage    int     `json:"Skonto1Tage"`
	Skonto1Prozent float64 `json:"Skonto1Prozent"`
	Skonto2Tage    int     `json:"Skonto2Tage"`
	Skonto2Prozent float64 `json:"Skonto2Prozent"`
}
