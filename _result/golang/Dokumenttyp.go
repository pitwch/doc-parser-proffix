package golang

type Dokumenttyp struct {
	DokumenttypNr string `json:DokumenttypNr`
	Bezeichnung   string `json:Bezeichnung`
	Inaktiv       bool   `json:Inaktiv`
	Parameter     string `json:Parameter`
}
