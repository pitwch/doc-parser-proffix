package golang

type Stammdaten struct {
	Anschrift    string `json:Anschrift`
	Berechtigung int    `json:Berechtigung`
	EMail        string `json:EMail`
	Fax          string `json:Fax`
	Firma        string `json:Firma`
	Homepage     string `json:Homepage`
	NoAutoLogin  bool   `json:NoAutoLogin`
	Ort          string `json:Ort`
	PLZ          string `json:PLZ`
	Sprache      int    `json:Sprache`
	Strasse      string `json:Strasse`
	Telefon      string `json:Telefon`
}
