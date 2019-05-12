package golang

type Teilleistung struct {
	TeilleistungNr           float64                  `json:TeilleistungNr`
	Anteil                   float64                  `json:Anteil`
	Bezeichnung              string                   `json:Bezeichnung`
	Gliederung               string                   `json:Gliederung`
	Teilleistungsgruppe      Teilleistungsgruppe      `json:Teilleistungsgruppe`
	Teilleistungstyp         Teilleistungstyp         `json:Teilleistungstyp`
	Teilleistungsuntergruppe Teilleistungsuntergruppe `json:Teilleistungsuntergruppe`
}
