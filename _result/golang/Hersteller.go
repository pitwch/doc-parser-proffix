package golang

type Hersteller struct {
	HerstellerNr float64 `json:HerstellerNr`
	EMail        string  `json:EMail`
	Fax          string  `json:Fax`
	Firma        string  `json:Firma`
	Homepage     string  `json:Homepage`
	Region       Region  `json:Region`
	Land         Land    `json:Land`
	Ort          string  `json:Ort`
	PLZ          string  `json:PLZ`
	Strasse      string  `json:Strasse`
	Telefon      string  `json:Telefon`
}
