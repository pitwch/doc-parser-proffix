package golang

type KonditionText struct {
	KonditionTextNr int       `json:"KonditionTextNr"`
	Bezeichnung     string    `json:"Bezeichnung"`
	Benutzer        Kondition `json:"Benutzer"`
	Modul           Sprache   `json:"Modul"`
}
