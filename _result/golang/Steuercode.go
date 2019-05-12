package golang

type Steuercode struct {
	SteuercodeNr float64 `json:SteuercodeNr`
	Bezeichnung  string  `json:Bezeichnung`
	Prozent      float64 `json:Prozent`
	InklusivMWST bool    `json:InklusivMWST`
	Konto        Konto   `json:Konto`
	GueltigVon   string  `json:GueltigVon`
	GueltigBis   string  `json:GueltigBis`
}