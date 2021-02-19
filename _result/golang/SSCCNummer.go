package golang

type SSCCNummer struct {
	SSCCNummerNr      int      `json:SSCCNummerNr`
	Artikel           Artikel  `json:Artikel`
	Dokument          Dokument `json:Dokument`
	Gewicht           float64  `json:Gewicht`
	GtinStufe         string   `json:GtinStufe`
	Haltbarkeitsdatum string   `json:Haltbarkeitsdatum`
	Menge             float64  `json:Menge`
	MengePalette      float64  `json:MengePalette`
	Paketnummer       string   `json:Paketnummer`
	PositionNr        int      `json:PositionNr`
	SSCCNummer        string   `json:SSCCNummer`
	SSCCOriginal      string   `json:SSCCOriginal`
	Tara              float64  `json:Tara`
	Verpackungen      int      `json:Verpackungen`
	ZuVerkaufenBis    string   `json:ZuVerkaufenBis`
}
