package golang

type Liste struct {
	ListeNr      int       `json:ListeNr`
	Name         string    `json:Name`
	Bezeichnung  string    `json:Bezeichnung`
	Art          string    `json:Art`
	Modul        string    `json:Modul`
	Datenbank    Datenbank `json:Datenbank`
	Sortierungen string    `json:Sortierungen`
	Selektionen  string    `json:Selektionen`
	Sortierungen string    `json:Sortierungen`
	Selektionen  string    `json:Selektionen`
}
