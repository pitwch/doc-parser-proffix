package golang

type Dokumenttypentext struct {
	DokumenttypentextNr int         `json:"DokumenttypentextNr"`
	Bezeichnung         string      `json:"Bezeichnung"`
	Dokumenttyp         Dokumenttyp `json:"Dokumenttyp"`
	Sprache             Sprache     `json:"Sprache"`
	Text                Text        `json:"Text"`
}
