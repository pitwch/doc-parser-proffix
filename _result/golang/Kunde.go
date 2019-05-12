package golang

type Kunde struct {
	KundeNr     int         `json:KundeNr`
	Adresse     Adresse     `json:Adresse`
	Prioritaet  Prioritaet  `json:Prioritaet`
	Vertragstyp Vertragstyp `json:Vertragstyp`
	Geloescht   bool        `json:Geloescht`
	Parameter   string      `json:Parameter`
}
