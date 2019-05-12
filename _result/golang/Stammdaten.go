package golang

type Stammdaten struct {
	Anschrift    string  `json:Anschrift`
	Berechtigung float64 `json:Berechtigung`
	EMail        string  `json:EMail`
	Fax          string  `json:Fax`
	Firma        string  `json:Firma`
	Homepage     string  `json:Homepage`
	NoAutoLogin  bool    `json:NoAutoLogin`
	Ort          string  `json:Ort`
	PLZ          string  `json:PLZ`
	Sprache      float64 `json:Sprache`
	Strasse      string  `json:Strasse`
	Telefon      string  `json:Telefon`
}
